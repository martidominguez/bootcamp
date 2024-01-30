package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigAppDefault is the default config for the application.
type ConfigAppDefault struct {
	// Router is the router implementation
	Router *chi.Mux
	// ServerAdress is the server address
	ServerAddress string
	// ConfigMySQL is the config for MySQL
	ConfigMySQL *mysql.Config
}

func NewApplicationDefaultDB(config *ConfigAppDefault) *ApplicationDefaultDB {
	// default config
	defaultConfig := &ConfigAppDefault{
		Router:        chi.NewRouter(),
		ServerAddress: ":8080",
		ConfigMySQL:   &mysql.Config{},
	}
	if config != nil {
		if config.Router != nil {
			defaultConfig.Router = config.Router
		}
		if config.ServerAddress != "" {
			defaultConfig.ServerAddress = config.ServerAddress
		}
		if config.ConfigMySQL != nil {
			defaultConfig.ConfigMySQL = config.ConfigMySQL
		}
	}

	return &ApplicationDefaultDB{
		router:        defaultConfig.Router,
		serverAddress: defaultConfig.ServerAddress,
		configMySQL:   defaultConfig.ConfigMySQL,
	}
}

// New returns a new instance of Application.
type ApplicationDefaultDB struct {
	// router is the router implementation.
	router *chi.Mux
	// serverAddress is the server address.
	serverAddress string
	// configMySQL is the MySQL configuration.
	configMySQL *mysql.Config
}

// SetUp sets up the application.
func (a *ApplicationDefaultDB) SetUp() (db *sql.DB, err error) {
	// dependencies
	// - database: open connection
	db, err = sql.Open("mysql", a.configMySQL.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	// - database: ping
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ping happened")
	// - repository
	rp := repository.NewRepositoryProductDatabase(db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	// routes
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - handler
	a.router.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefaultDB) Run() (err error) {
	err = http.ListenAndServe(a.serverAddress, a.router)
	return
}
