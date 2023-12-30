package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Structs
type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	servers         []Server
}

// Interfaces
type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

// Functions
func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func newSimpleServer(address string) *SimpleServer {
	serverUrl, err := url.Parse(address)
	handleError(err)

	proxy := httputil.NewSingleHostReverseProxy(serverUrl)

	return &SimpleServer{
		Addr:  address,
		Proxy: proxy,
	}
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		servers:         servers,
	}
}

// Struct methods
func (lb *LoadBalancer) getNextAvailableServer() Server {
	if len(lb.servers) == 0 {
		fmt.Println("No available servers")
		os.Exit(1)
	}

	server := lb.servers[lb.RoundRobinCount%len(lb.servers)]
	for server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.servers[lb.RoundRobinCount%len(lb.servers)]
	}
	lb.RoundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}

func (s *SimpleServer) IsAlive() bool {
	return true
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}

// Main function
func main() {
	servers := []Server{
		newSimpleServer("http://www.facebook.com"),
		newSimpleServer("http://www.bing.com"),
		newSimpleServer("http://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}

	http.Handle("/", http.HandlerFunc(handleRedirect))

	fmt.Printf("Serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
