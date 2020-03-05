package application

import (
	"XlsForXMLHttp/config"
	"XlsForXMLHttp/funchttp"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kardianos/service"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type App struct {
	log      *zap.Logger
	funcHttp *funchttp.HttpApp
	conf     *config.Config
	Router    *mux.Router
}

func AppNew(log *zap.Logger, funcHttp *funchttp.HttpApp, conf *config.Config) *App {
	s := &App{
		conf:   conf,
		log:    log,
		funcHttp: funcHttp,
		Router: mux.NewRouter(),
	}
	s.Router.HandleFunc("/healthcheck/", healthCheckHandler)
	return s
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	return
}

func (p *App) Start(s service.Service) error {
	if service.Interactive() {
		log.Println("XlsForXMLHttp running in terminal. It is not correct run a programm")
	} else {
		log.Println("XlsForXMLHttp running under service manager.")

	}
	go p.RegisterHandlers(p.Router)
	return nil
}


func (p *App) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/payorder/files/test", p.funcHttp.Test).Methods("GET")
	router.HandleFunc("/payorder/files/zapsib", p.funcHttp.GetOnlyZB).Methods("GET")
	router.HandleFunc("/payorder/files/sber", p.funcHttp.GetOnly).Methods("GET")
	router.HandleFunc("/payorder/backup", p.funcHttp.ToArch).Methods("GET")
	router.HandleFunc("/payorder/clear", p.funcHttp.ClearDir).Methods("GET")
	p.log.Info(fmt.Sprintf("Server running, port: %v", p.conf.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", p.conf.Port), nil)
	if err != nil {
		panic(err)
	}
}

func (p *App) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
