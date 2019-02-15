package main

import (
        "fmt"
        "net/http"
        "log"
        "encoding/json"
)
func getGithubConfig() string {

	resp, err := http.Get("https://raw.githubusercontent.com/g278701988/go-getting-started/master/config.txt")
	if err != nil {
		log.Printf("Get err%v", err)
		return string("")
	}
	var jsData []Jsdata
	if err := json.NewDecoder(resp.Body).Decode(&jsData); nil != err {
		log.Printf("Decode err%v", err)
		return string("")
	}

	return fmt.Sprintf("%v:%v\n(update every half hour)", jsData[0].Key, jsData[0].Value)
}
func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
