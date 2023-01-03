package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/acool-kaz/parser-service-server/internal/config"
	"github.com/acool-kaz/parser-service-server/internal/service"
)

type Handler struct {
	cfg     *config.Config
	service *service.Service
}

func InitHandler(cfg *config.Config, service *service.Service) *Handler {
	log.Println("init http handler")
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	log.Println("init routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, this is root handler go to /parser - to parse")
	})

	mux.HandleFunc("/parser", func(w http.ResponseWriter, r *http.Request) {
		err := h.service.Post.Parse(r.Context(), h.cfg.Parser.ParseUrl, h.cfg.Parser.MaxPage)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}
		fmt.Fprintln(w, "OK")
	})

	return mux
}
