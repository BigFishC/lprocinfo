package src

import (
	"io/ioutil"
	"log"
	"strings"
)

//Loadavg -
type Loadavg struct {
	Onemin     string `json:"onemin"`
	Fivemin    string `json:"fivemin"`
	Fifteen    string `json:"fifteen"`
	Perprocess string `json:"perprocess"`
	Runprocess string `json:"runprocess"`
}

//GetLoad -
func GetLoad(path string) *Loadavg {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	newloadavgslice := strings.Split(string(r), " ")
	return &Loadavg{
		Onemin:     newloadavgslice[0],
		Fivemin:    newloadavgslice[1],
		Fifteen:    newloadavgslice[2],
		Perprocess: newloadavgslice[3],
		Runprocess: newloadavgslice[4],
	}
}
