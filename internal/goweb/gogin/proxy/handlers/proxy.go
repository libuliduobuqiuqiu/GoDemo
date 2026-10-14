package handlers

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func parseProxyTarget(raw string) (*url.URL, error) {
	target, err := url.Parse(raw)
	if err != nil || target == nil || (target.Scheme != "http" && target.Scheme != "https") || target.Hostname() == "" || target.User != nil || target.Fragment != "" {
		return nil, errors.New("proxy_pass must be an absolute HTTP or HTTPS URL without credentials or fragment")
	}
	return target, nil
}

func proxyOrigin(target *url.URL) string {
	return (&url.URL{Scheme: target.Scheme, Host: target.Host}).String()
}

// NewProxyHandler permits only origins explicitly selected at server startup.
func NewProxyHandler(allowedTargets []string) (gin.HandlerFunc, error) {
	allowed := make(map[string]bool, len(allowedTargets))
	for _, raw := range allowedTargets {
		target, err := parseProxyTarget(raw)
		if err != nil || target.Path != "" || target.RawQuery != "" {
			return nil, errors.New("--proxy-target must be an HTTP or HTTPS origin")
		}
		allowed[proxyOrigin(target)] = true
	}

	return func(ctx *gin.Context) {
		target, err := parseProxyTarget(ctx.Query("proxy_pass"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !allowed[proxyOrigin(target)] {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "proxy target is not allowed"})
			return
		}

		proxy := &httputil.ReverseProxy{
			Rewrite: func(req *httputil.ProxyRequest) {
				req.SetURL(target)
				req.Out.URL.Path = target.Path
				req.Out.URL.RawPath = target.RawPath
				req.Out.URL.RawQuery = target.RawQuery
			},
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				http.Error(w, "upstream request failed", http.StatusBadGateway)
			},
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}, nil
}
