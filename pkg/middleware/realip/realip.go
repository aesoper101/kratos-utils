package realip

import (
	"github.com/aesoper101/kratos-utils/pkg/utils"
	"net/http"
)

func Handler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if rip := utils.RealIP(r); rip != "" {
			r.RemoteAddr = rip
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
