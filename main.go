package main

import (
	"log"
	"net/http"
	"net"
	"net/http/httputil"
	"time"
	"os"
)

func main() {
	http.HandleFunc("/", ReverseProxy())
	println("ready")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func ReverseProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: func(network, addr string) (net.Conn, error) {
				return getConnection(req)
			},
			TLSHandshakeTimeout: 10 * time.Second,
		}
		(&httputil.ReverseProxy{
			Director: func(req *http.Request) {
				println("forwarding to http://"+os.Getenv("DST_IP") + ":" + os.Getenv("DST_PORT"))
				req.URL.Scheme = "http"
				req.URL.Host = "/" + req.RequestURI
			},
			Transport: transport,
		}).ServeHTTP(w, req)
	}
}

func getConnection(req *http.Request) (net.Conn, error) {
	println("forwarding to http://"+os.Getenv("DST_IP") + ":" + os.Getenv("DST_PORT"))
	return net.Dial("tcp", os.Getenv("DST_IP") + ":" + os.Getenv("DST_PORT"))
}
