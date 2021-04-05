package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"context"
)

type Server interface {
	Start() error
	Stop() error
	Addr() string
}

type Options struct {
	Logger     *log.Logger
	ShowQuery  bool
	ShowBody   bool
	ShowHeader bool
}

type server struct {
	instance *http.Server
	listener net.Listener
	ops      Options
	log      *log.Logger
}

func Create(ops Options) Server {
	if ops.Logger == nil {
		ops.Logger = log.Default()
	}

	ip, _ := net.LookupIP("localhost")
	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:0", ip[0]))

	srvr := &http.Server{
		Handler:      CreateHandler(ops),
		WriteTimeout: time.Second * 2,
		ReadTimeout:  time.Second * 2,
	}

	return &server{ops: ops, log: ops.Logger, instance: srvr, listener: listener}
}

func (s *server) Start() (err error) {

	go func() {
		err = s.instance.Serve(s.listener)
	}()

	return err
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return s.instance.Shutdown(ctx)
}

func (s *server) Addr() string {
	return s.listener.Addr().String()
}
