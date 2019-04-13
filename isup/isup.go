package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

var state struct {
	sync.RWMutex
	yes bool
}

var (
	listenAddr = flag.String("http", "localhost:8080", "Listen address")
	pollUrl    = flag.String("pollUrl", "http://localhost:8081/hello", "Polling address")
	pollPeriod = flag.Duration("poll", 5*time.Second, "Polling period")
)

var tml = template.Must(template.New("root").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>Is server {{.URL}} running?</h2>
	<h1>
		{{if .YES}}
			YES :)
		{{else}}
			NO :(
		{{end}}
	</h1>
</center><body><html>
`))

func main() {
	flag.Parse()
	go poll(*pollPeriod)
	http.HandleFunc("/is-up", handler)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	state.RLock()
	data := struct {
		URL string
		YES bool
	}{
		*pollUrl,
		state.yes,
	}
	state.RUnlock()
	_ = tml.Execute(writer, data)
}

func poll(period time.Duration) {
	for true {
		state.Lock()
		state.yes = isUp()
		state.Unlock()
		time.Sleep(period)
	}
}

func isUp() bool {
	_, err := http.Get(*pollUrl)
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
