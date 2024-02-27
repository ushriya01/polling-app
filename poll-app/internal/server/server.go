package server

import (
	"context"
	"fmt"
	"net/http"
	"poll-app/internal/handlers"
	"poll-app/internal/middleware"

	"github.com/julienschmidt/httprouter"
)

// Server represents the HTTP server
type Server struct {
	Addr string
	ctx  context.Context
}

// NewServer initializes a new HTTP server
func NewServer(addr string, ctx context.Context) *Server {
	return &Server{Addr: addr, ctx: ctx}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	router := s.SetupRoutes()
	fmt.Printf("Server listening on %s\n", s.Addr)
	return http.ListenAndServe(s.Addr, router)
}

func (s *Server) SetupRoutes() http.Handler {
	router := httprouter.New()

	router.GET("/api/polls", adaptHandler(handlers.ListPollsHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.POST("/api/polls", adaptHandler(handlers.CreatePollHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.PUT("/api/polls/:pollID", adaptHandler(handlers.UpdatePollHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.GET("/api/polls/:pollID", adaptHandler(handlers.GetPollHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.DELETE("/api/polls/:pollID", adaptHandler(handlers.DeletePollHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.POST("/api/polls/vote", adaptHandler(handlers.VoteHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.GET("/api/polls/:pollID/options/:optionID/votes", adaptHandler(handlers.GetVotesHandler, s.ctx, middleware.ParseTokenMiddleware))
	router.POST("/api/signup", adaptHandler(handlers.SignupHandler, s.ctx))
	router.POST("/api/signin", adaptHandler(handlers.SigninHandler, s.ctx))
	router.POST("/api/clear-data", adaptHandler(handlers.ClearDataHandler, s.ctx, middleware.ParseTokenMiddleware))
	return router
}

// adaptHandler adapts a httprouter.Handle to httprouter.Handle
func adaptHandler(handler httprouter.Handle, ctx context.Context, middlewares ...func(httprouter.Handle) httprouter.Handle) httprouter.Handle {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler(w, r.WithContext(ctx), ps)
	}
}
