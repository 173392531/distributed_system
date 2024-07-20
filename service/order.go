package service

import (
	"encoding/json"
	"net/http"

	"github.com/173392531/distributed_system/discovery"
)

type OrderService struct {
	address         string
	discoveryClient *discovery.ConsulClient
}

func NewOrderService(address string, discoveryClient *discovery.ConsulClient) *OrderService {
	return &OrderService{address, discoveryClient}
}

func (s *OrderService) GetOrders(w http.ResponseWriter, r *http.Request) {
	// TODO: 从数据库中获取订单列表
	orders := []string{"Order1", "Order2", "Order3"}
	json.NewEncoder(w).Encode(orders)
}
