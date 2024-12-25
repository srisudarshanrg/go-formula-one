package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/go-setup-template/server/config"
	"github.com/srisudarshanrg/go-setup-template/server/database"
	"github.com/srisudarshanrg/go-setup-template/server/functions"
	"github.com/srisudarshanrg/go-setup-template/server/handlers"
	"github.com/srisudarshanrg/go-setup-template/server/render"
	"github.com/srisudarshanrg/go-setup-template/server/validations"
)

const portNumber = ":{put_your_port_number_here}"

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
	defer db.Close()

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
		AppConfig: appConfig,
	}
	handlers.RepositoryAccesshandlers(handlerRepo)

	// app config access
	functions.AppConfigAccessFunctions(appConfig)
	validations.AppConfigAccessValidations(appConfig)

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

	mux.Use(SessionLoadAndSave)

	mux.Get("/", handlers.AppConfig.Home)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func SessionLoadAndSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
