package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImage(doc)
	return
}

func countWordsAndImage(n *html.Node) (words, images int) {
	return
}
