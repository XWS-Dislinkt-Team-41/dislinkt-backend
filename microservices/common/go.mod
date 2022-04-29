module github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common

go 1.16

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/user => ../common/user

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
