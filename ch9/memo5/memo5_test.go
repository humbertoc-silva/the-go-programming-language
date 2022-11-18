package memo5

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var incomingURLs []string = []string{
	"https://go.dev/",
	"https://pkg.go.dev/",
	"https://go.dev/play/",
	"http://gopl.io",
	"https://go.dev/",
	"https://pkg.go.dev/",
	"https://go.dev/play/",
	"http://gopl.io",
}

func Test(t *testing.T) {
	m := New(httpGetBody)
	for _, url := range incomingURLs {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	m.Close()
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	m.Close()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
