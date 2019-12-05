package auth

import (
	"log"
	"net/http"

	"github.com/micro/micro/plugin"
	"github.com/myproject-0722/mn-hosted/lib/token"
)

// JWTAuthWrapper JWT鉴权Wrapper
func JWTAuthWrapper(token *token.Token) plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth plugin received: " + r.URL.Path)
			// TODO 从配置中心动态获取白名单URL
			if r.URL.Path == "/order/Order/AlipayNotify" || r.URL.Path == "/alipay" || r.URL.Path == "/user/User/SignIn" || r.URL.Path == "/user/User/SignUp" || r.URL.Path == "/user/User/MailCode" || r.URL.Path == "/user/User/Reset" || r.URL.Path == "/user/test" {
				h.ServeHTTP(w, r)
				return
			}

			tokenstr := r.Header.Get("Authorization")
			userFromToken, e := token.Decode(tokenstr)

			if e != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			log.Println("User Name : ", userFromToken.UserName)
			r.Header.Set("X-Example-Username", userFromToken.UserName)
			h.ServeHTTP(w, r)
		})
	}
}
