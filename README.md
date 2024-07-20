# distributed_system
distributed system by go

go-distributed-system/
├── config/
│   └── config.go
├── discovery/
│   └── consul.go
├── gateway/
│   └── gateway.go
├── service/
│   ├── user.go
│   └── order.go
└── main.go


<!-- 按以下顺序执行 -->
go mod tidy
go build

go run main.go