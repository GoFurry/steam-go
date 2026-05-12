# alpha-3 审计处理记录

本轮 `alpha-3` 审计项已经全部处理完成，以下内容保留为修复记录，便于后续回溯。

## 已修复

- `[Fixed] P1-001` 代理指标中的 `ProxyURL` 现在默认脱敏，不再暴露代理用户名和密码。
- `[Fixed] P1-002` 新增 `steam.RedactSensitiveURL(...)`，并在文档中明确要求不要直接记录带 `key` / `access_token` 的原始请求 URL。
- `[Fixed] P2-001` 新增 `WithSafeDefaults()`，为真实外部调用提供保守的默认重试和限流预设。
- `[Fixed] P2-002` 新增 `WithHealthCheckedAPIKeys(...)`，支持对反复触发 `401/429` 的 API key 进行临时冷却，避免坏 key 持续污染轮转。
- `[Fixed] P2-003` 默认 client 现在会克隆 `http.DefaultTransport`，避免多个 SDK client 共享同一个全局 transport 实例。
- `[Fixed] P2-004` OpenID `claimed_id` 现在仅接受 `https://steamcommunity.com/openid/id/...`。
- `[Fixed] P3-001` OpenID 示例成功页不再回显 `state`。
- `[Fixed] P3-002` CI 已补充 `go vet ./...`、`go test -race ./...`、`staticcheck`、`govulncheck`。
- `[Fixed] P3-003` OpenID 已补充边界测试和 fuzz 测试入口，覆盖 `claimed_id` / `return_to` 的畸形输入场景。
- `[Fixed] P3-004` README 和中文文档已明确区分 `OfficialAPI` 与 `PublicStorePage` 的使用边界，不再示范把 typed 官方 API 调用错误地包进商店页流量类别。

## 说明

- 本文件当前不再保留未处理项。
- 下一次代码审计开始前，可以再次清空本文件，保持审计结果单次闭环。
