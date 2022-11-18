package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("hit the page")
//		next.ServeHTTP(w, r)
//	})
//}

// csrf-middleware 添加csrf protection 给每个会话
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// 对每一个请求加载并保存会话
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
