package error

import "net/http"

func Validate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.ContentLength < 5 {
			// a valid request is passed on to next handler
			next.ServeHTTP(w, req)
		} else {
			// otherwise, respond with an error
			http.Error(w, "Bad request", 400)
		}
	}
	return http.HandlerFunc(fn)
}
