package helper

import (
	"fmt"
	"log"
	"os"
	"strings"
	"tools/utils"
)

type DivideInfo struct {
	Yamls  *[]utils.YamlInfo
	input  string
	output string
}

func NewDivideInfo(yamls *[]utils.YamlInfo, input, output string) *DivideInfo {
	return &DivideInfo{
		Yamls:  yamls,
		output: output,
		input:  input,
	}
}

func (d *DivideInfo) RunSort() error {

	dir, err := os.ReadDir(d.input)
	if err != nil {
		return err
	}

	var listFiles = []string{}
	for _, file := range dir {
		listFiles = append(listFiles, file.Name())
	}

	total := 0
	for _, continent := range *d.Yamls {

		dirContinent := fmt.Sprintf("%s/%s", d.output, continent.Name)
		if err := utils.EnsureDir(dirContinent); err != nil {
			return err
		}
		diff := match(continent.Countries, listFiles)
		log.Printf("%d 	Number of %s", len(diff), continent.Name)
		MoveFile(diff, d.input, dirContinent)
		total += len(diff)
	}
	log.Printf("total countries sorted: %d", total)

	return nil
}

func match(countries, files []string) []string {
	c := make(map[string]bool)
	var match []string

	for _, item := range countries {
		c[item] = true
	}

	for _, item := range files {
		uperCItem := strings.ToUpper(strings.Split(item, ".")[0])
		if _, ok := c[uperCItem]; ok {
			match = append(match, item)
		}
	}

	return match
}

func MoveFile(files []string, path, folder string) {
	for _, file := range files {
		source := fmt.Sprintf("%s/%s", path, file)
		destination := fmt.Sprintf("%s/%s", folder, file)
		err := os.Rename(source, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
}
