package main

import (
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
)

const Name = "simple"

func init() {
	http.RegisterHttpFilterFactoryAndConfigParser(Name, ConfigFactory, http.NullParser)
}

func ConfigFactory(c interface{}, callbacks api.FilterCallbackHandler) api.StreamFilter {
	return &filter{
		callbacks: callbacks,
	}
}

func main() {}
