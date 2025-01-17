package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/go-formula-one/server/config"
	"github.com/srisudarshanrg/go-formula-one/server/database"
	"github.com/srisudarshanrg/go-formula-one/server/functions"
	"github.com/srisudarshanrg/go-formula-one/server/handlers"
	"github.com/srisudarshanrg/go-formula-one/server/render"
	"github.com/srisudarshanrg/go-formula-one/server/validations"
)

const portNumber = ":2221"

var session *scs.SessionManager
var appConfig config.AppConfig

func main() {
	// session
	session = scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// database handlers
	db, err := database.CreateDatabaseConn()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// template cache handlers
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not create template cache: ", err)
	}

	// app config handlers
	appConfig.TemplateCache = templateCache
	appConfig.ProjectCompleted = false
	appConfig.UseTemplateCache = appConfig.ProjectCompleted
	appConfig.Database = db
	appConfig.Session = session

	// handlers repository
	handlerRepo := handlers.HandlerRepository{
		AppConfig: &appConfig,
	}
	handlers.RepositoryAccesshandlers(&handlerRepo)

	// app config access
	functions.AppConfigAccessFunctions(&appConfig)
	validations.AppConfigAccessValidations(&appConfig)

	// routes
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("server running on port number", portNumber)
	server.ListenAndServe()
}

func routes() http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(SessionLoadAndSave)

	// routes for get requests
	mux.Get("/", handlers.Repository.Home)
	mux.Get("/drivers", handlers.Repository.Drivers)

	// routes for post requests

	// routes for static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func SessionLoadAndSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
