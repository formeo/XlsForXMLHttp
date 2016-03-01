// XlsForOra project main.go
package main

/**/
import (
	"github.com/formeo/XlsForOra/funchttp"
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"net/http"
	"os"
)

var logger service.Logger
var install bool
var uninstall bool

type program struct{}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		log.Println("XlsForOra running in terminal. It is not correct run a programm")
		//return nil
	} else {
		log.Println("XlsForOra running under service manager.")

	}
	go p.run()
	return nil
}
func (p *program) run() {
	http.HandleFunc("/payorder/files/test", funchttp.Test)
	http.HandleFunc("/payorder/files/zapsib", funchttp.GetOnlyZB)
	http.HandleFunc("/payorder/files/sber", funchttp.GetOnly)
	http.HandleFunc("/payorder/backup", funchttp.ToArch)
	http.HandleFunc("/payorder/clear", funchttp.ClearDir)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

//CommandLineParse Парсит командную строку
func CommandLineParse() {
	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, "\t-install\tinstall as service\n")
		fmt.Fprintf(os.Stderr, "\t-uninstall\tuninstall as service\n")

	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(os.Stderr, "Error:%s\n", e)
			flag.Usage()
			os.Exit(1)
		}
	}()

	flag.BoolVar(&install, "install", false, "installing as service")
	flag.BoolVar(&uninstall, "uninstall", false, "uninstalling as service")
	help := flag.Bool("h", false, "usage")
	flag.BoolVar(help, "help", false, "usage")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
}

func run(s service.Service) {

	if install {
		s.Install()
		return
	}

	if uninstall {
		s.Uninstall()
		return
	}

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
func main() {
	CommandLineParse()
	svcConfig := &service.Config{
		Name:        "GoXlsForOraService",
		DisplayName: "Go Service XlsForOra",
		Description: "This is an XlsForOra Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	run(s)

}
