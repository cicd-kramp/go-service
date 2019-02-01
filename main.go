package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "expvar"

    httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
    "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func main() {
    tracer.Start(tracer.WithServiceName("go-dog"))
    defer tracer.Stop()

    mux := httptrace.NewServeMux() // init the http tracer
    mux.HandleFunc("/", sayhelloName) // use the tracer to handle the urls
    mux.HandleFunc("/dog", sayHello) // use the tracer to handle the urls
    mux.Handle("/debug/vars", http.DefaultServeMux) // use the tracer to handle the urls

    err := http.ListenAndServe(":9090", mux) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
