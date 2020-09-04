package httpserver

import (
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

// RetrieveMetadata takes token from cookie and put it into context
func RetrieveMetadata(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// retrive token
		var token string

		//--
		log.Printf("headers: %+v\n", r.Header)
		//--

		authorization := r.Header.Get("Authorization")
		if authorization != "" {
			splits := strings.Split(authorization, ":")
			if len(splits) > 0 {
				token = strings.TrimSpace(splits[0])
			}
		} else {
			cookie, err := r.Cookie("token_user")
			if err == nil {
				token = cookie.Value
			}
		}

		// retrive ip address
		var ip string
		ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
		if len(ips) > 0 && ips[0] != "" {
			ip = ips[0]
		} else {
			ipAddr := r.RemoteAddr
			var err error
			ip, _, err = net.SplitHostPort(ipAddr)
			if err != nil {
				log.Println(err)
			}
		}

		// retrive company id
		var companyID string
		cookie, err := r.Cookie("company_id")
		if err == nil {
			companyID = cookie.Value
		}

		// retrive company UI language
		var uiLang string
		cookie, err = r.Cookie("selected_lang")
		if err == nil {
			uiLang = cookie.Value
		}

		ctx := metadata.NewOutgoingContext(r.Context(), metadata.Pairs(
			"token", token,
			"ip", ip,
			"company_id", companyID,
			"ui_lang", uiLang,
			"http_user_agent", r.UserAgent(),
		))

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
