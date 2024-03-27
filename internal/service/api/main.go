package api

import (
	"github.com/tandesko/cards_validator/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
	"net"
	"net/http"
)

type service struct {
	cfg      config.Config
	log      *logan.Entry
	listener net.Listener
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		cfg:      cfg,
		log:      cfg.Log(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
