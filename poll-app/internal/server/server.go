package server

import (
	"context"
	"fmt"
	"net/http"
	"poll-app/internal/handlers"

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

	router.GET("/api/polls", adaptHandler(handlers.ListPollsHandler, s.ctx))
	router.POST("/api/polls", adaptHandler(handlers.CreatePollHandler, s.ctx))
	router.PUT("/api/polls/:pollID", adaptHandler(handlers.UpdatePollHandler, s.ctx))
	router.GET("/api/polls/:pollID", adaptHandler(handlers.GetPollHandler, s.ctx))
	router.DELETE("/api/polls/:pollID", adaptHandler(handlers.DeletePollHandler, s.ctx))
	router.POST("/api/polls/vote", adaptHandler(handlers.VoteHandler, s.ctx))
	router.GET("/api/polls/:pollID/options/:optionID/votes", adaptHandler(handlers.GetVotesHandler, s.ctx))
	router.POST("/api/signup", adaptHandler(handlers.SignupHandler, s.ctx))
	router.POST("/api/signin", adaptHandler(handlers.SigninHandler, s.ctx))

	return router
}

// adaptHandler adapts a httprouter.Handle to httprouter.Handle
func adaptHandler(handler httprouter.Handle, ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler(w, r.WithContext(ctx), ps)
	}
}
