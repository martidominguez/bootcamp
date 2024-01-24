package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	loader := loader.NewLoaderTicketCSV(a.dbFile)
	ticketsMap, err := loader.Load()
	if err != nil {
		fmt.Println("the error is: ", err)
	}
	// repository
	rp := repository.NewRepositoryTicketMap(ticketsMap, 0)
	// service
	sv := service.NewServiceTicketDefault(rp) // check this later
	// handler
	hd := handler.NewHandlerTicketsDefault(sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})
	(*a).rt.Get("/tickets", hd.GetTotalTickets())

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
