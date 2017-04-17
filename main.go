package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
)


type Site struct {
	Name string `json:"name"`
	URL  []string `json:"url"`
	Path string `json:"path"`
}

//type Config struct {
//	Sites []Site `json:"sites"`
//}
var sites map[string]Site
func main() {
	//var c Config
	//file, _ := os.Open("config.json")
	file, _ := ioutil.ReadFile("config.json")
	// /err := json.NewDecoder(file).Decode(&c)

	err := json.Unmarshal(file, &sites)
	if err != nil {
		panic(err)
	}

	for _, v := range sites {
		fmt.Printf("%s:%s\t%s\n", v.Name, v.URL, v.Path)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ServeItUp)
	http.ListenAndServe(":80", router)
}

func ServeItUp(w http.ResponseWriter, r *http.Request) {
	site := fmt.Sprintf("%s", getSite(sites,"localhost").Path)
	fmt.Println(site)
	http.ServeFile(w,r,site)


}

func getSite(sites map[string]Site, url string) Site {
	for _, value := range sites {
		for _, checkUrl := range value.URL {
			if checkUrl == url {
				return value
			}
		}
	}
	return Site{}
}