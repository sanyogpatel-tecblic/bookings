package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanyogpatel-tecblic/bookings/internal/config"
	"github.com/sanyogpatel-tecblic/bookings/internal/driver"
	"github.com/sanyogpatel-tecblic/bookings/internal/handlers"
	"github.com/sanyogpatel-tecblic/bookings/internal/helpers"
	"github.com/sanyogpatel-tecblic/bookings/internal/models"
	"github.com/sanyogpatel-tecblic/bookings/internal/render"
)

const portNumber = ":8090"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.Mailchan)

	listetnformail()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})
	gob.Register(models.Room{})
	gob.Register(models.User{})
	gob.Register(models.RoomRestriction{})

	Mailchan := make(chan models.MailData)
	app.Mailchan = Mailchan

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	//connect to database
	log.Println("connecting to database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=root")

	if err != nil {
		log.Println("Couldn't connect to the database")
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}
	log.Println("Connected to database!")

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	render.NewRenderer(&app)
	return db, nil
}
