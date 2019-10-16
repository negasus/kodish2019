package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Hello, Kodish")
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	mux.HandleFunc("/info", showInfo)
	http.ListenAndServe(":80", mux)
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, Developer!")
}

func showInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Your IP address is %s\n", req.RemoteAddr)
	fmt.Fprintf(w, "Your UserAgent: %s\n", req.Header.Get("user-agent"))
}
