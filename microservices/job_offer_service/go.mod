module github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/job_offer_service

go 1.16

replace github.com/XWS-Dislinkt-Team-41/dislinkt-backend/microservices/common => ../common

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
