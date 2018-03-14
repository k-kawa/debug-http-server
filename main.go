package main

import (
    "fmt"
    "time"
    "net/http"
    "log"
    "encoding/json"
    "io/ioutil"
    "flag"
)

type LogFormat struct {
	Method string
	URL string
	Proto string
    Header http.Header
    Body string
}

var port string

func init() {
    flag.StringVar(&port, "port", "9000", "Port number to listen")
    flag.Parse()
}

func LoggingHandler(w http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Printf("error: %s", err.Error())
    }
    defer req.Body.Close()
    
    l := LogFormat{
        Method: req.Method,
        URL: req.URL.String(),
        Proto: req.Proto,
        Header: req.Header,
        Body: string(body),
    }

    s, err := json.Marshal(l)
    if err != nil {
        log.Printf("error: %s", err.Error())
    }

    log.Println(string(s))
}

func main() {
    srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(LoggingHandler),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	 	MaxHeaderBytes: 1 << 20,
	}
    log.Fatal(srv.ListenAndServe())
}
