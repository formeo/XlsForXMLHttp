package main

import (
	"flag"
	"fmt"
	"github.com/formeo/XlsForXMLHttp/application"
	"github.com/formeo/XlsForXMLHttp/config"
	"github.com/formeo/XlsForXMLHttp/funchttp"
	"github.com/formeo/XlsForXMLHttp/logger"
	"github.com/getsentry/sentry-go"
	"github.com/kardianos/service"
	"os"
)

var install bool
var uninstall bool

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

func run(s service.Service) error {
	var err error
	if install {
		if err = s.Install(); err != nil {
			return err
		}
		return nil
	}

	if uninstall {
		if err = s.Uninstall(); err != nil {
			return err
		}
		return nil
	}

	if err = s.Run(); err != nil {
		return err
	}
	return nil
}
func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	err = logger.InitSentry(conf)
	if err != nil {
		panic(err)
	}
	log, err := logger.NewLogger(conf)
	if err != nil {
		sentry.CaptureException(err)
	}

	httpFunc := funchttp.NewHttpApp(log, conf)
	CommandLineParse()
	svcConfig := &service.Config{
		Name:        "GoXlsForXMLHTTPService",
		DisplayName: "Go Service XlsForXMLHttp",
		Description: "This is an XlsForXMLHttp Go service.",
	}

	prg := application.AppNew(log, httpFunc, conf)
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = run(s); err != nil {
		log.Fatal(err.Error())
	}
}
