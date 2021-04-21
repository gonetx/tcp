# gonetx/tcp
<p align="center">
  <a href="https://pkg.go.dev/github.com/gonetx/tcp">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
  </a>
  <a href="https://goreportcard.com/report/github.com/gonetx/tcp">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
  </a>
  <a href="https://codecov.io/gh/gonetx/tcp">
    <img src="https://codecov.io/gh/gonetx/tcp/branch/main/graph/badge.svg?token=05UN78X11O"/>
  </a>
  <a href="https://github.com/gonetx/tcp/actions?query=workflow%3ASecurity">
    <img src="https://img.shields.io/github/workflow/status/gonetx/tcp/Security?label=%F0%9F%94%91%20gosec&style=flat&color=75C46B">
  </a>
  <a href="https://github.com/gonetx/tcp/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/workflow/status/gonetx/tcp/Test?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B">
  </a>
</p>

`tcp` is a package providing utility functions for tcp.

# Detect
Detect checks if a tcp address is connectable with specific timeout.

```go
ok, err := tcp.Detect("www.google.com:443", time.Second)
fmt.Println(ok, err)
```
