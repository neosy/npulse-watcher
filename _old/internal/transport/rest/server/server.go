package server

import (
	"fmt"
	"log"

	uc "git.n-hub.ru/neosy/npulse-watcher/internal/usecase"
	"github.com/valyala/fasthttp"
)

type HTTPServer struct {
	Compress bool
	uc       *uc.UseCase
}

// Создания объекта для fastHTTP сервера
func New(uc *uc.UseCase) (*HTTPServer, error) {
	s := &HTTPServer{}

	s.uc = uc

	return s, nil
}

// Запуск fastHTTP сервера
func (s *HTTPServer) RunServer(address string, port int) {
	h := s.HandlerMain

	if s.Compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", address, port), h); err != nil {
		log.Panicf("Error in ListenAndServe: %v", err)
	}

}
