package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"tools/helper"
	"tools/utils"
)

func main() {

	// log config
	logName := "app.log"
	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// pars flag
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\tExample: sort-countries -c conf.yaml -d countries\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	ymlFile := flag.String("c", "", "yaml file path")
	collection := flag.String("d", "", "countries path")

	flag.Parse()

	if flag.NFlag() != 2 {
		flag.Usage()
		return
	}

	conf := utils.Newconfig(*ymlFile)
	conf.Initial()
	if err := conf.GetConf(); err != nil {
		log.Fatal(err)
	}

	absPath, err := filepath.Abs(*collection)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("%s/%s", conf.BasePath, "continents")

	sort := helper.NewDivideInfo(conf.Yamls, absPath, output)
	if err := sort.RunSort(); err != nil {
		log.Fatal(err)
	}

}
