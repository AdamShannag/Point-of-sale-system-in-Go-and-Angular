PROTO_PD_DIRECTORY=pb
BROKER_PD_DIRECTORY=../broker-service/proto/auth

proto:
	mkdir -p ${BROKER_PD_DIRECTORY}
	mkdir -p ${PROTO_PD_DIRECTORY}
	protoc --proto_path=proto --go_out=${BROKER_PD_DIRECTORY} --go_out=${PROTO_PD_DIRECTORY} --go_opt=paths=source_relative \
	--go-grpc_out=${BROKER_PD_DIRECTORY} --go-grpc_out=${PROTO_PD_DIRECTORY} --go-grpc_opt=paths=source_relative \
	proto/*.proto

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis
