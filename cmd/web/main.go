package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/lcxznpy/bookings/pkg/config"
	"github.com/lcxznpy/bookings/pkg/handles"
	"github.com/lcxznpy/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// 让中间件也可以使用这些变量
var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false
	//创建会话
	session = scs.New()
	session.Lifetime = 24 * time.Hour              //该会话持续24h
	session.Cookie.Persist = true                  //保存cookie
	session.Cookie.SameSite = http.SameSiteLaxMode //这个cookie对于相同的站点都适应
	session.Cookie.Secure = app.InProduction       //cookie 保护
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handles.NewRepo(&app)
	handles.NewHandles(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handles.Repo.Home)
	//http.HandleFunc("/about", handles.Repo.About)

	fmt.Printf(fmt.Sprintf("starting on port %s", portNumber))
	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
