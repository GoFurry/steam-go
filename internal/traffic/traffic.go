package traffic

import "context"

// Class identifies one request traffic category inside the SDK.
type Class string

const (
	ClassOfficialAPI     Class = "official_api"
	ClassPublicStorePage Class = "public_store_page"
)

type classContextKey struct{}

// WithClass attaches one traffic class to a context.
func WithClass(ctx context.Context, class Class) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, classContextKey{}, NormalizeClass(class))
}

// ClassFromContext resolves one traffic class from context.
func ClassFromContext(ctx context.Context) (Class, bool) {
	if ctx == nil {
		return "", false
	}
	class, ok := ctx.Value(classContextKey{}).(Class)
	if !ok {
		return "", false
	}
	return NormalizeClass(class), true
}

// NormalizeClass coerces empty or unknown values back to the default official API class.
func NormalizeClass(class Class) Class {
	switch class {
	case ClassPublicStorePage:
		return ClassPublicStorePage
	case ClassOfficialAPI:
		fallthrough
	default:
		return ClassOfficialAPI
	}
}
