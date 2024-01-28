package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/config"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/handlers"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function 
func main() {
	

	// change this to true when in production 
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour 
	session.Cookie.Persist = true 
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	fmt.Printf("Starting application at %s\n", portNumber)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
}