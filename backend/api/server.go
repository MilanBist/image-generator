package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/image-generator/config"
	"github.com/image-generator/internal/middlewares"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Server struct{
	Router 	*chi.Mux
	Db		*pgxpool.Pool
	Config 	*config.Config
}

func(srv *Server) setupRoutes(){
	srv.Router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success in getting response."))
	})

	srv.Router.Route("/api", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Success"))
		})

		// setup for the login and the registration
		r.Post("/register", srv.HandleRegister)
		r.Post("/login", srv.HandleLogin)
		r.Get("/refreshToken", srv.HandleNewAccessToken)


		// create the protected handlers
		r.Group(func(r chi.Router) {
			// add the middlewares here
			r.Use(middlewares.LoggingMiddleware)
			r.Use(middlewares.AuthenticationMiddleware)

			// after making these handlers protected now use certain things here


		})
	})
}


func NewServer(db *pgxpool.Pool, cfg *config.Config) (*Server, error){
	r := chi.NewRouter()


	// use the middlewares here
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// if the time is more than 60 s
	// use the middleware for the time if more than 60 abort
	r.Use(middleware.Timeout(60 *time.Second))

	// use the cors here
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
  }))


	// now create a server involving all of the components required
	srv := &Server{
		Router: r,
		Db: db,
		Config: cfg,
	}


	srv.setupRoutes()

	return srv,nil
}
