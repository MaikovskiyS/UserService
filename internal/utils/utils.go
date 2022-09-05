package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadJson(r http.Request) (v uint, err error) {
	var dat map[string]*uint
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println(dat)
	fmt.Printf("%T", dat["id"])
	v = *dat["id"]
	return v, nil
}
