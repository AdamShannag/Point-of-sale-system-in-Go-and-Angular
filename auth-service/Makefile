PROTO_PD_DIRECTORY=internal/grpc/proto/auth
BROKER_PD_DIRECTORY=../broker-service/internal/grpc/proto/auth
DB_URL=postgresql://pos:pos@localhost:5432/pos?sslmode=disable

proto: proto_broker
	mkdir -p -m777 ${PROTO_PD_DIRECTORY}
	protoc --proto_path=proto  --go_out=${PROTO_PD_DIRECTORY} --go_opt=paths=source_relative \
  --go-grpc_out=${PROTO_PD_DIRECTORY} --go-grpc_opt=paths=source_relative \
  proto/*.proto


proto_broker:
	mkdir -p ${BROKER_PD_DIRECTORY}
	protoc --proto_path=proto --go_out=${BROKER_PD_DIRECTORY} --go_opt=paths=source_relative \
  --go-grpc_out=${BROKER_PD_DIRECTORY} --go-grpc_opt=paths=source_relative \
  proto/*.proto

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down


sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: sqlc  server  proto proto_broker migrateup migratedown
