package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type RSS struct {
	XMLName     xml.Name `xml:"rss"`
	AtomNS      string   `xml:"xmlns:atom,attr"`
	Version     string   `xml:"version,attr"`
	Channel     Channel  `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Generator     string `xml:"generator"`
	Language      string `xml:"language"`
	LastBuildDate string `xml:"lastBuildDate"`
	AtomLink      AtomLink `xml:"atom:link"`
	Items         []Item `xml:"item"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}


func fetchFeed(url string) ([]Item, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RSS
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Channel.Items, nil
}

// THIS IS A DEVELOPMENT TEST FUNCTION AND WILL BE DELETED

// func testFetchFeed(w http.ResponseWriter, r *http.Request) {
// 	data, err := fetchFeed("https://blog.boot.dev/index.xml")
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respBody := struct {
// 		Data []Item `json:"data"`
		
// 	}{
// 		Data: data,
// 	}

// 	respondWithJSON(w, http.StatusOK, respBody)
// }