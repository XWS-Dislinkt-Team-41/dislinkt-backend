module github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/notification_service

go 1.18

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common => ../common

require (
	github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common v1.0.0
	github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service v0.0.0-20220508211809-c225be35129b
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/grpc v1.47.0
)

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208 // indirect
	golang.org/x/sys v0.0.0-20220429121018-84afa8d3f7b3 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220630174209-ad1d48641aa7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
