package main

import (
	"strings"

	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
)

var globeRS string

type filter struct {
	api.PassThroughStreamFilter
	callbacks api.FilterCallbackHandler
}

func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	result := globeRS
	tmp, ok := header.Get("cookie")
	if ok {
		result = result + "; " + tmp
	}
	header.Set("cookie", result)
	return api.Continue
}

func (f *filter) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	//info := f.callbacks.StreamInfo()
	//ignoreF := func(s string, ok bool) string {
	//	return s
	//}
	//f.callbacks.Log(api.Warn, fmt.Sprintf("downStream[%s||%s],upStream[%s||%s],name[%s||%s],",
	//	info.DownstreamRemoteAddress(), info.DownstreamLocalAddress(),
	//	ignoreF(info.UpstreamRemoteAddress()), ignoreF(info.UpstreamLocalAddress()),
	//	ignoreF(info.UpstreamClusterName()), ignoreF(info.VirtualClusterName())))
	if globeRS == "" {
		get, _ := header.Get("set-cookie")
		if strings.Contains(get, "global-session-cookie") {
			split := strings.Split(get, `;`)
			globeRS = split[0]
		}
	}
	return api.Continue
}
