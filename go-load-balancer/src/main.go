package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type SimpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// create simpleServer instance
func createSimpleServer(address string) *SimpleServer {
	serverUrl, err := url.Parse(address)
	handleErr(err)

	return &SimpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}

}

//create LoadBalancer intance

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0, //init
		servers:         servers,
	}
}

func handleErr(err error) {

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

func main() {

	//have the servers created 

	servers := []Server{
		createSimpleServer("http://www.facebook.com")

	}

}
