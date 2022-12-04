setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/database.yml.example internal/config/db/database.yml
	@echo " --- Done Setup and generate configuration --- "

rest:swagger
	@go run main.go
	@#go run cmd/server/restful/main.go

swagger:
	@swag init

build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main main.go