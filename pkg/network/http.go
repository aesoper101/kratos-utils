package network

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http/pprof"
)

func RegisterMetrics(srv *http.Server, metrics *confpb.MetricsSettings) {
	if metrics != nil && metrics.GetEnabled() {
		path := "/metrics"
		if handle := metrics.GetHandle(); len(handle) > 0 {
			path = handle
		}
		srv.Handle(path, promhttp.Handler())
	}
}

func RegisterSwagger(srv *http.Server, swagger *confpb.SwaggerSettings) {
	if swagger != nil && swagger.GetEnabled() {
		openAPIhandler := openapiv2.NewHandler()
		prefix := "/q/"
		if handle := swagger.GetHandle(); len(handle) > 0 {
			prefix = handle
		}
		srv.HandlePrefix(prefix, openAPIhandler)
	}
}

func RegisterPprof(s *http.Server, pprofCfg *confpb.PprofSettings) {
	if pprofCfg != nil && pprofCfg.GetEnabled() {
		prefix := "/metrics"
		if handle := pprofCfg.GetPrefix(); len(handle) > 0 {
			prefix = handle
		}
		s.HandleFunc(prefix+"/pprof", pprof.Index)
		s.HandleFunc(prefix+"/pprof/cmdline", pprof.Cmdline)
		s.HandleFunc(prefix+"/pprof/profile", pprof.Profile)
		s.HandleFunc(prefix+"/pprof/symbol", pprof.Symbol)
		s.HandleFunc(prefix+"/pprof/trace", pprof.Trace)
		s.HandleFunc(prefix+"/allocs", pprof.Handler("allocs").ServeHTTP)
		s.HandleFunc(prefix+"/block", pprof.Handler("block").ServeHTTP)
		s.HandleFunc(prefix+"/goroutine", pprof.Handler("goroutine").ServeHTTP)
		s.HandleFunc(prefix+"/heap", pprof.Handler("heap").ServeHTTP)
		s.HandleFunc(prefix+"/mutex", pprof.Handler("mutex").ServeHTTP)
		s.HandleFunc(prefix+"/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	}
}
