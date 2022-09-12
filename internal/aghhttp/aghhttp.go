// Package aghhttp provides some common methods to work with HTTP.
package aghhttp

import (
	"fmt"
	"io"
	"net/http"

	"github.com/AdguardTeam/AdGuardHome/internal/version"
	"github.com/AdguardTeam/golibs/log"
)

// RegisterFunc is the function that sets the handler to handle the URL for the
// method.
//
// TODO(e.burkov, a.garipov):  Get rid of it.
type RegisterFunc func(method, url string, handler http.HandlerFunc)

// OK responds with word OK.
func OK(w http.ResponseWriter) {
	if _, err := io.WriteString(w, "OK\n"); err != nil {
		log.Error("couldn't write body: %s", err)
	}
}

// Error writes formatted message to w and also logs it.
func Error(r *http.Request, w http.ResponseWriter, code int, format string, args ...any) {
	text := fmt.Sprintf(format, args...)
	log.Error("%s %s: %s", r.Method, r.URL, text)
	http.Error(w, text, code)
}

// UserAgent returns the ID of the service as a User-Agent string.  It can also
// be used as the value of the Server HTTP header.
func UserAgent() (ua string) {
	return fmt.Sprintf("AdGuardDNS/%s", version.Version())
}
