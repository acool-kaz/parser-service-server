package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/acool-kaz/parser-service-server/internal/config"
	grpcParserHandler "github.com/acool-kaz/parser-service-server/internal/delivery/grpc/parser"
	httpHandler "github.com/acool-kaz/parser-service-server/internal/delivery/http"
	"github.com/acool-kaz/parser-service-server/internal/repository"
	"github.com/acool-kaz/parser-service-server/internal/service"
	parser_pb "github.com/acool-kaz/parser-service-server/pkg/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type app struct {
	cfg *config.Config

	db *sql.DB

	httpServer  *http.Server
	httpHandler *httpHandler.Handler

	grpcServer        *grpc.Server
	grpcParserHandler *grpcParserHandler.ParserHandler
}

func InitApp(cfg *config.Config) (*app, error) {
	log.Println("init app")
	db, err := repository.InitDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("init app: %w", err)
	}

	repo := repository.InitRepository(db)
	service := service.InitService(repo)

	httpHandler := httpHandler.InitHandler(cfg, service)

	grpcParserHandler := grpcParserHandler.InitParserHandler(cfg, service)

	return &app{
		cfg:               cfg,
		db:                db,
		httpHandler:       httpHandler,
		grpcParserHandler: grpcParserHandler,
	}, nil
}

func (a *app) RunApp() {
	log.Println("run app")
	go func() {
		if err := a.startHTTP(); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Println("http server started on", a.cfg.Http.Port)

	go func() {
		if err := a.startGRPC(); err != nil {
			log.Println(err)
			return
		}
	}()
	log.Println("grpc started on", a.cfg.Grpc.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
		return
	}
	a.grpcServer.GracefulStop()
	log.Println("grpc: Server closed")

	if err := a.db.Close(); err != nil {
		log.Println(err)
	} else {
		log.Println("db closed")
	}
}

func (a *app) startGRPC() error {
	listen, err := net.Listen(a.cfg.Grpc.Type, fmt.Sprintf("%s:%s", a.cfg.Grpc.Host, a.cfg.Grpc.Port))
	if err != nil {
		return err
	}
	defer listen.Close()

	opt := []grpc.ServerOption{}

	a.grpcServer = grpc.NewServer(opt...)

	parser_pb.RegisterParserServiceServer(a.grpcServer, a.grpcParserHandler)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listen)
}

func (a *app) startHTTP() error {
	router := a.httpHandler.InitRoutes()

	a.httpServer = &http.Server{
		Handler:      router,
		Addr:         ":" + a.cfg.Http.Port,
		ReadTimeout:  time.Second * time.Duration(a.cfg.Http.Read),
		WriteTimeout: time.Second * time.Duration(a.cfg.Http.Write),
	}

	return a.httpServer.ListenAndServe()
}
