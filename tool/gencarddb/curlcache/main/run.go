package main

import (
	"log"
	"main/cachedhttp"
	"main/exporter"

	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	setupGin("/cachedir").Run(":80")
}

func setupGin(cacheDir string) *gin.Engine {
	reg := exporter.NewRegistry()
	engine := gin.Default()
	engine.GET("/", root(cacheDir))
	engine.GET("/metrics", prometheusHandler(reg))
	pprof.Register(engine)
	return engine
}

func prometheusHandler(reg *prometheus.Registry) gin.HandlerFunc {
    h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

var (
	CurlcacheHitHeaderKey = "Curlcache-Hit"
	CurlcacheHitHeaderValue = map[bool]string{true: "true", false: "false"}
)

func root(cacheDir string) func(c *gin.Context) {
	return func(c *gin.Context)	{
		exporter.ReqCounter.With(prometheus.Labels{"code": "all"}).Inc()

		encodedUrl := c.Query("q")
		encodedReferer := c.Query("referer")

		if encodedUrl == "" {
			c.String(400, "q is required query string")
			return
		}

		// analyse query parameter
		planeUrl, err := base64url.Decode(encodedUrl)
		if err != nil {
			log.Println(err)
			c.String(400, "can not parse query")
			return
		}
		log.Printf("target url is %s", planeUrl)
		
		planeReferer, err := base64url.Decode(encodedReferer)
		if err != nil {
			log.Println(err)
			c.String(400, "can not parse query")
			return
		}
		log.Printf("referer is %s", planeReferer)

		data, hit, err := cachedhttp.Get(cacheDir, string(planeUrl), string(planeReferer))
		if err != nil {
			log.Println(err)
			c.String(400, "query url is not valid")
			c.String(400, "\nif you want access example.com")
			c.String(400, "\nq=http%3A%2F%2Fexample.com%2F")
			return
		}

		exporter.ReqCounter.With(prometheus.Labels{"code": "200"}).Inc()

		c.Header(CurlcacheHitHeaderKey, CurlcacheHitHeaderValue[hit])
		c.String(200, string(data))
	}
}
