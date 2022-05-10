# oauth

基于 go 标准库和 html 实现的 OAuth 集成示例.

目前支持

- [github]()
- [gitee]()
- ...

## FAQ

### Q1:三方的user信息拿到后，要怎么绑定和处理?

### Q2:多个三方OAuth对接时同一个callback实现如何区分是哪个三方过来的?

一种方案是，每一个三方OAuth对接时，都是用不同的callback地址和实现来做区分。

那么如果我们要实现仅使用同一个callback地址和实现时如何区分是哪一个三方过来的，以便区分后续导致调用哪个三方的TokenAPI?

参考

- [构建OAuth应用程序 - Github Docs](https://docs.github.com/cn/developers/apps/building-oauth-apps/authorizing-oauth-apps)
- [Gitee OAuth 文档](https://gitee.com/api/v5/oauth_doc#/)
