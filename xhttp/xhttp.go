package xhttp

import (
	"encoding/json"
	"net/http"

	"github.com/mesg-foundation/core/x/xstrings"
)

// ResponseJSON responses to http request with given status code and json data.
func ResponseJSON(w http.ResponseWriter, status int, data interface{}) error {
	bdata, err := json.Marshal(data)
	if err != nil {
		status = http.StatusInternalServerError
		bdata, _ = json.Marshal(map[string]string{"error": http.StatusText(status)})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bdata)
	return err
}

// ParseJSON parses json payload from request's payload into o.
func ParseJSON(r *http.Request, o interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(o)
}

// HasHeader checks if http request has header with given values. if values are nil or len(zero),
// it'll only check for presence of header.
func HasHeader(r *http.Request, header string, values []string) bool {
	vs, ok := r.Header[header]
	if !ok {
		return false
	}
	for _, v := range values {
		if !xstrings.SliceContains(vs, v) {
			return false
		}
	}
	return true
}
