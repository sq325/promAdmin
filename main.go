package main

import (
	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/sq325/kitComplement/instrumentation"
	"github.com/sq325/promAdmin/pkg/prom"
	"github.com/sq325/promAdmin/pkg/proxy"
)

var (
	consulAddr = pflag.String("consul.addr", "", "eg: localhost:8500, required")
	port       = pflag.StringP("port", "p", "8080", "server port")

	version = pflag.BoolP("version", "v", false, "show version info")
)

const (
	_service     = "promAdmin"
	_version     = "v0.0.0"
	_versionInfo = "init"
)

var (
	buildTime      string
	buildGoVersion string
	author         string
)

func init() {
	pflag.Parse()
}

func main() {
	if *version {
		fmt.Println(_service, _version)
		fmt.Println("build time:", buildTime)
		fmt.Println("go version:", buildGoVersion)
		fmt.Println("author:", author)
		return
	}
	if *consulAddr == "" {
		panic("consul.addr is required")
	}

	promSvc := prom.NewPromSvc(*consulAddr)
	httpProxy := proxy.NewProxy()

	mux := gin.Default()
	mux.GET(
		"/services",
		instrumentation.GinHandlerFunc(
			"GET",
			"/services",
			prom.MakePromSvcEndpoint(promSvc),
			prom.DecodeInstancesRequest,
			prom.EncodeResponse,
		),
	)

	mux.Use(static.Serve("/dist", static.LocalFile("./dist", false)))
	mux.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})
	mux.Any("/proxy", httpProxy.Request)

	mux.Run(":" + *port)
}
