package main

import (
	"ecom/internal/products"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// mount - create all the endpoints
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting - avoids DDOS
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)                 // recover from crashes
	r.Use(middleware.Timeout(60 * time.Second)) // 60 second timeout

	// Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProducts)

	return r
}

// run - start server and do graceful shutdown
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at address: %s\n", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	// db driver
}

type config struct {
	addr string // address of the server
	db   dbConfig
}
type dbConfig struct {
	dsn string // a string that has the username password
}
