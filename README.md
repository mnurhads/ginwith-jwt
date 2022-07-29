# GIN with JWT - API Authorizen JWT

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://travis-ci.org/gin-gonic/gin.svg)](https://travis-ci.org/gin-gonic/gin)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)
[![Build Status](https://secure.travis-ci.org/auth0/node-jsonwebtoken.svg?branch=master)](http://travis-ci.org/auth0/node-jsonwebtoken)

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter). If you need performance and good productivity, you will love Gin.


## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.13+ is required**), then you can use the below Go command to install Gin.

```sh
$ go get -u github.com/gin-gonic/gin
```

2. Import it in your code:

```go
import "github.com/gin-gonic/gin"
```

3. (Optional) Import `net/http`. This is required for example if using constants such as `http.StatusOK`.

```go
import "net/http"
```

## Quick start

```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

```
# run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
$ go run example.go
```
