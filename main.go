package main

import (
	"log"
	"net/http"

	"github.com/173392531/distributed_system/config"
	"github.com/173392531/distributed_system/discovery"
	"github.com/173392531/distributed_system/gateway"
	"github.com/173392531/distributed_system/service"
	"github.com/gorilla/mux"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 创建服务发现客户端
	discoveryClient := discovery.NewConsulClient(cfg.ConsulAddress)

	// 创建用户服务
	userService := service.NewUserService(cfg.UserServiceAddress, discoveryClient)

	// 创建订单服务
	orderService := service.NewOrderService(cfg.OrderServiceAddress, discoveryClient)

	// 创建API网关
	gateway := gateway.NewGateway(discoveryClient)

	// 注册路由
	router := mux.NewRouter()
	router.HandleFunc("/users", userService.GetUsers).Methods("GET")
	router.HandleFunc("/orders", orderService.GetOrders).Methods("GET")
	router.HandleFunc("/api/{service}/{path}", gateway.ProxyRequest)

	// 启动HTTP服务器
	log.Printf("Starting server on %s", cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, router))
}
