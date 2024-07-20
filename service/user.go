package service

import (
	"encoding/json"
	"net/http"

	"github.com/173392531/distributed_system/discovery"
)

type UserService struct {
	address         string
	discoveryClient *discovery.ConsulClient
}

func NewUserService(address string, discoveryClient *discovery.ConsulClient) *UserService {
	return &UserService{address, discoveryClient}
}

func (s *UserService) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: 从数据库中获取用户列表
	users := []string{"Alice", "Bob", "Charlie"}
	json.NewEncoder(w).Encode(users)
}
