package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var cfg *Config

//Config тип для конфига указывается папка с файлами и папка для бэкапа
type Config struct {
	PathToFiles        string `json:"PathToFiles"`
	PathToBackupFolder string `json:"PathToBackupFolder"`
	PathToClearDir     string `json:"PathToClearDir"`
	file               string
}

//Current Функция возвращает текущий конфиг
func Current() *Config {
	if cfg == nil {
		fileconfig := filepath.Join(filepath.Dir(os.Args[0]), "config.json")
		if cfg == nil {
			var err error
			cfg, err = New(fileconfig)
			cfg.file = fileconfig
			if err != nil {
				panic(err)
			}
		}
	}
	return cfg
}

//New Функция создает новый конфиг
func New(filename string) (result *Config, err error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	if e = json.Unmarshal(file, &result); e != nil {
		return nil, e
	}
	return result, nil
}
