package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IPaddress struct {
	PublicIP string `json:"publicip"`
}
type ipaddress []IPaddress

func json_func(w http.ResponseWriter, r *http.Request) {
	m := getip()
	// var i = 1
	// data, _ := json.Unmarshal(m)

	// fmt.Print(data)
	// IP := IPaddress{PublicIP: m}
	// bytes, _ := json.Marshal(IP)
	// fmt.Print("JSON", string(bytes))
	ipaddresss := ipaddress{

		IPaddress{PublicIP: m},
	}

	// fmt.Print(i)

	json.NewEncoder(w).Encode(ipaddresss)

}

type IP struct {
	Query string
}

// var bb []IPaddress

func main() {
	handleRequest()
}

func home(w http.ResponseWriter, r *http.Request) {
	m := getip()
	fmt.Fprint(w, m)
}
func handleRequest() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/home1", json_func)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func getip() string {

	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}
