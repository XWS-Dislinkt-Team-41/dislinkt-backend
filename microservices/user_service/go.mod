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
	go.mongodb.org/mongo-driver v1.8.4
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common v1.0.0
