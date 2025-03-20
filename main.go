package main

import (
	"github.com/alanwade2001/go-sepa-engine-data/repository"
	"github.com/alanwade2001/go-sepa-engine-export/internal/handler"
	"github.com/alanwade2001/go-sepa-engine-export/internal/service"

	inf "github.com/alanwade2001/go-sepa-infra"
)

type App struct {
	Infra   *inf.Infra
	Manager *repository.Manager
	Service *service.Export
	Handler *handler.Export
}

func NewApp() *App {
	infra := inf.NewInfra()
	manager := repository.NewManager(infra.Persist)
	service := service.NewExport(manager)
	handler := handler.NewExport(service, infra.Router)

	app := &App{
		Infra:   infra,
		Manager: manager,
		Service: service,
		Handler: handler,
	}

	return app
}

func (a *App) Run() {
	a.Infra.RunWithTLS()
}

func main() {
	app := NewApp()

	app.Run()

}
