package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/config"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/handlers"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/models"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/internal/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function 
func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}
	
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

func run() error {

	//What am I going to to put in the session 
	gob.Register(models.Reservation{})

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
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)


	return nil 
}