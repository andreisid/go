package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type News struct {
	Locations []Location `xml:"url"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprint(l.Loc)
}

func main() {
	resp, err := http.Get("https://www.washingtonpost.com/news-politics-sitemap.xml")
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
	resp.Body.Close()
	var n News
	xml.Unmarshal(bodyBytes, &n)

	fmt.Println(n.Locations)
}
