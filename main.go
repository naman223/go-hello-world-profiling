package main
import (
        "fmt"
        "net/http"
        "os"
        "net/http/pprof"
        _ "net/http/pprof"
)

const (
        port = ":9999"
)
var calls = 0

func sayHello(w http.ResponseWriter, r *http.Request) {
        h, _ := os.Hostname()
        calls++
		fmt.Fprintf(w, "Hello! You have called me %d times from %s!", calls, h)
}


func main() {
		fmt.Println("Started http Server",port)
        r := http.NewServeMux()
        r.HandleFunc("/", sayHello)

        // Register pprof handlers
        r.HandleFunc("/debug/pprof/", pprof.Index)
        r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
        r.HandleFunc("/debug/pprof/profile", pprof.Profile)
        r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
        r.HandleFunc("/debug/pprof/trace", pprof.Trace)
        http.ListenAndServe(port, r)
}