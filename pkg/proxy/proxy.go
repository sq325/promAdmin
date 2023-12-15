package proxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type Proxy interface {
	Request(c *gin.Context)
}

type proxy struct{}

func NewProxy() Proxy {
	return &proxy{}
}

func (p *proxy) Request(c *gin.Context) {
	host := c.Query("host")   // host or host:port
	path := c.Query("path")   // path (relative paths may omit leading slash)
	query := c.Query("query") // without ?
	if query != "" {
		path += "?" + query
	}
	u := p.getUrl(host, path)
	request, _ := http.NewRequest(c.Request.Method, u.String(), c.Request.Body)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.String(resp.StatusCode, string(body))
}

func (p *proxy) getUrl(host, path string) *url.URL {
	if !strings.HasPrefix("http://", host) {
		host = "http://" + host
	}
	u, _ := url.Parse(host)

	if path != "" {
		u = u.JoinPath(path)
	}
	return u
}
