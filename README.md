<div align="center">
  <a href="https://encore.dev" alt="encore"><img width="189px" src="https://encore.dev/assets/img/logo.svg"></a>
  <h3><a href="https://encore.dev">Encore – The Backend Development Engine</a></h3>
</div>

# Encore Runtime API

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/encore.dev.svg)](https://pkg.go.dev/encore.dev) [![Slack](https://img.shields.io/badge/chat-on%20slack-blue?style=flat-square&logo=slack)](https://encore.dev/slack) ![MPL-2 License](https://img.shields.io/github/license/encoredev/encore.dev?style=flat-square)

</div>

This repository contains the [Encore's](https://encore.dev) Runtime API contract, which Encore applications are built against.

```bash
go get -u encore.dev
```

When compiling your application the [encore compiler](https://github.com/encoredev/encore) will automatically provide an
implementation of this API. [The default implementation can be found here](https://github.com/encoredev/encore/tree/main/runtime).
