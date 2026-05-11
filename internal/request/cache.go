package request

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/GoFurry/steam-go/internal/traffic"
)

type CacheRuntime interface {
	lookup(req *http.Request, now time.Time) cacheLookup
	store(req *http.Request, resp *http.Response, body []byte, now time.Time)
	refresh(lookup cacheLookup, resp *http.Response, now time.Time) ([]byte, bool)
}

type cacheLookup struct {
	key          string
	body         []byte
	etag         string
	lastModified string
	fresh        bool
	found        bool
}

type memoryCacheRuntime struct {
	ttl       time.Duration
	cookieJar http.CookieJar

	mu      sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	body         []byte
	etag         string
	lastModified string
	storedAt     time.Time
	expiresAt    time.Time
}

type cacheKeyPayload struct {
	Method         string   `json:"method"`
	URL            string   `json:"url"`
	SessionKey     string   `json:"session_key,omitempty"`
	AcceptLanguage string   `json:"accept_language,omitempty"`
	ExplicitCookie string   `json:"explicit_cookie,omitempty"`
	JarCookies     []string `json:"jar_cookies,omitempty"`
}

func NewMemoryCacheRuntime(ttl time.Duration, jar http.CookieJar) CacheRuntime {
	if ttl <= 0 {
		return nil
	}
	return &memoryCacheRuntime{
		ttl:       ttl,
		cookieJar: jar,
		entries:   make(map[string]cacheEntry),
	}
}

func (c *memoryCacheRuntime) lookup(req *http.Request, now time.Time) cacheLookup {
	if c == nil {
		return cacheLookup{}
	}
	key, ok := c.cacheKey(req)
	if !ok {
		return cacheLookup{}
	}

	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	if !ok {
		return cacheLookup{}
	}

	return cacheLookup{
		key:          key,
		body:         cloneBytes(entry.body),
		etag:         entry.etag,
		lastModified: entry.lastModified,
		fresh:        !entry.expiresAt.Before(now),
		found:        true,
	}
}

func (c *memoryCacheRuntime) store(req *http.Request, resp *http.Response, body []byte, now time.Time) {
	if c == nil || req == nil || resp == nil || req.Method != http.MethodGet {
		return
	}
	key, ok := c.cacheKey(req)
	if !ok {
		return
	}

	c.mu.Lock()
	c.entries[key] = cacheEntry{
		body:         cloneBytes(body),
		etag:         strings.TrimSpace(resp.Header.Get("ETag")),
		lastModified: strings.TrimSpace(resp.Header.Get("Last-Modified")),
		storedAt:     now,
		expiresAt:    now.Add(c.ttl),
	}
	c.mu.Unlock()
}

func (c *memoryCacheRuntime) refresh(lookup cacheLookup, resp *http.Response, now time.Time) ([]byte, bool) {
	if c == nil || !lookup.found || lookup.key == "" {
		return nil, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[lookup.key]
	if !ok {
		return nil, false
	}

	entry.storedAt = now
	entry.expiresAt = now.Add(c.ttl)
	if resp != nil {
		if etag := strings.TrimSpace(resp.Header.Get("ETag")); etag != "" {
			entry.etag = etag
		}
		if lastModified := strings.TrimSpace(resp.Header.Get("Last-Modified")); lastModified != "" {
			entry.lastModified = lastModified
		}
	}
	c.entries[lookup.key] = entry
	return cloneBytes(entry.body), true
}

func (c *memoryCacheRuntime) cacheKey(req *http.Request) (string, bool) {
	if req == nil || req.URL == nil || req.Method != http.MethodGet {
		return "", false
	}

	payload := cacheKeyPayload{
		Method:         req.Method,
		URL:            req.URL.String(),
		AcceptLanguage: req.Header.Get("Accept-Language"),
		ExplicitCookie: req.Header.Get("Cookie"),
	}
	if sessionKey, ok := traffic.RequestSessionKeyFromContext(req.Context()); ok {
		payload.SessionKey = sessionKey
	}
	if c.cookieJar != nil {
		payload.JarCookies = normalizedCookies(c.cookieJar.Cookies(req.URL))
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", false
	}
	return string(encoded), true
}

func normalizedCookies(cookies []*http.Cookie) []string {
	if len(cookies) == 0 {
		return nil
	}

	normalized := make([]string, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie == nil {
			continue
		}
		normalized = append(normalized, cookie.Name+"="+cookie.Value)
	}
	sort.Strings(normalized)
	if len(normalized) == 0 {
		return nil
	}
	return normalized
}

func cloneBytes(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	cloned := make([]byte, len(src))
	copy(cloned, src)
	return cloned
}

func applyConditionalCacheHeaders(req *http.Request, lookup cacheLookup) {
	if req == nil || !lookup.found || lookup.fresh {
		return
	}
	if req.Header.Get("If-None-Match") == "" && lookup.etag != "" {
		req.Header.Set("If-None-Match", lookup.etag)
	}
	if req.Header.Get("If-Modified-Since") == "" && lookup.lastModified != "" {
		req.Header.Set("If-Modified-Since", lookup.lastModified)
	}
}

func cacheLookupAllowsConditionalRequest(lookup cacheLookup) bool {
	return lookup.found && !lookup.fresh && (lookup.etag != "" || lookup.lastModified != "")
}

func requestCacheable(req *http.Request) bool {
	return req != nil && req.Method == http.MethodGet && req.URL != nil
}
