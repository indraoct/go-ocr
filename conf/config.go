package conf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Configuration struct {
	PathFile 			 string		`split_words:"true" default:"/tmp"`
	RootUrl 			 string		`split_words:"true" default:"go_ocr"`
	Port                 string   	`default:":8080"`
}

var Config Configuration


func InitConfig() {
	err := envconfig.Process("go-ocr", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Loaded config %#v", Config)
}