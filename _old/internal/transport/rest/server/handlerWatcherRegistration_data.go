package server

import (
	"log"

	"git.n-hub.ru/neosy/npulse-watcher/internal/models"
	"github.com/go-playground/validator"
)

type WatcherRegRequest struct {
	models.WatcherRegRequest
}

func (data *WatcherRegRequest) validate() (err error) {
	validate := validator.New()
	if err = validate.Struct(data); err != nil {
		log.Printf("Validate JSON\n")
		errs := err.(validator.ValidationErrors)
		for _, field := range errs {
			log.Printf("field %s: %s\n", field.Field(), field.Tag())
		}
	}
	return err
}
