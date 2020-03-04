package application

import (
	"github.com/formeo/XlsForXMLHttp/funchttp"
	"github.com/kardianos/service"
	"log"
	"net/http"
)

type App struct{}

func AppNew() *App{
	return &App{}
}

func (p *App) Start(s service.Service) error {
	if service.Interactive() {
		log.Println("XlsForOra running in terminal. It is not correct run a programm")
	} else {
		log.Println("XlsForOra running under service manager.")

	}
	go p.run()
	return nil
}
func (p *App) run() {
	http.HandleFunc("/payorder/files/test", funchttp.Test)
	http.HandleFunc("/payorder/files/zapsib", funchttp.GetOnlyZB)
	http.HandleFunc("/payorder/files/sber", funchttp.GetOnly)
	http.HandleFunc("/payorder/backup", funchttp.ToArch)
	http.HandleFunc("/payorder/clear", funchttp.ClearDir)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
func (p *App) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}