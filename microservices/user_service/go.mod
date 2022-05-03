module github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service

go 1.16

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common => ../common

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/domain => ../user_service/domain

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common/proto/user_service => ../common/proto/user_service

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/application => ../user_service/application

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/infrastructure/api => ../user_service/infrastructure/api

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/infrastructure/persistence => ../user_service/infrastructure/persistence

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/startup/config => ../user_service/startup/config

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/user_service/startup => ../user_service/startup

require (
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220429121018-84afa8d3f7b3 // indirect
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common v1.0.0
	go.mongodb.org/mongo-driver v1.9.0
	google.golang.org/grpc v1.46.0
)
