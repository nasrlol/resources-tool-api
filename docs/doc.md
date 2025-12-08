### API

The API provides both REST and Websocket endpoints. Each with their own purposes.


```postman

POST("user/login")
POST("user/register")

```
> The software provides other websocket enpoints that will deliver the live system resources when required.
> System software receives a UPDATE/POST request from the host device, which then gets "forwarded to the receiving user"

Structure

```bash
├── LICENSE
├── README.md
├── cmd
│   └── main.go
├── docs
│   └── doc.md
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── data
│   │   │   ├── models
│   │   │   │   ├── cpu.go
│   │   │   │   ├── device.go
│   │   │   │   ├── disk.go
│   │   │   │   ├── ram.go
│   │   │   │   ├── token.go
│   │   │   │   └── user.go
│   │   │   ├── rest
│   │   │   │   ├── device_handler.go
│   │   │   │   ├── disk_handler.go
│   │   │   │   ├── ram_handler.go
│   │   │   │   └── user_handler.go
│   │   │   ├── services
│   │   │   │   ├── cpu_service.go
│   │   │   │   ├── ram_service.go
│   │   │   │   └── token_service.go
│   │   │   └── ws
│   │   │       ├── cpu_handler.go
│   │   │       └── ram_handler.go
│   │   └── routes
│   │       └── routes.go
│   ├── cache
│   │   └── redis.go
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── db.go
│   ├── migrations
│   │   ├── 20250503171210_session.go
│   │   ├── 20250624124146_user.go
│   │   ├── 20250625001518_device.go
│   │   ├── 20250625005811_cpu.go
│   │   ├── 20250625010423_ram.go
│   │   ├── 20250625010714_disk.go
│   │   └── 20250802215820_token.go
│   └── tests
│       └── user_test.go
├── key
└── pkg
    ├── scripts
    │   └── goose
    │       └── goose-synf
    └── utils
        └── logger
            └── logger.go

```
