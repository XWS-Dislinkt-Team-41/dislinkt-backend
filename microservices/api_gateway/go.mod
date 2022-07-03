module github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/api_gateway

go 1.18

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common => ../common

require (
	github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/auth_service v0.0.0-20220508211809-c225be35129b
	github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	google.golang.org/grpc v1.46.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	go.mongodb.org/mongo-driver v1.9.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220429170224-98d788798c3e // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
