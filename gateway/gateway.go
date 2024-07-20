package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/173392531/distributed_system/discovery"
	"github.com/gorilla/mux"
)

type Gateway struct {
	discoveryClient *discovery.ConsulClient
}

func NewGateway(discoveryClient *discovery.ConsulClient) *Gateway {
	return &Gateway{discoveryClient}
}

func (g *Gateway) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceName := vars["service"]
	path := vars["path"]

	services, err := g.discoveryClient.GetService(serviceName)
	if err != nil || len(services) == 0 {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	service := services[0]
	targetURL, _ := url.Parse("http://" + service.Service.Address + "/" + path)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)
}
