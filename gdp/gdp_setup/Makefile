build_broker:
	cd ../brokerService && go mod tidy && env GOOS=linux CGO_ENABLED=0 go build -o brokerApp ./cmd/api
build_auth:
	cd ../authenticationService && go mod tidy && env GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd/api
build_inventory:
	cd ../inventoryService && go mod tidy && env GOOS=linux CGO_ENABLED=0 go build -o inventoryApp ./cmd/api

start_broker:
	cd ../brokerService && docker run -t --rm --name broker-service -p 8080:8080 broker-app

start_auth:
	cd ../authenticationService && docker run -t --rm --name auth-service -p 8081:8081 auth-app

start_inventory:
	cd ../inventoryService && docker run -t --rm --name inventory-service -p 8082:8082 inventory-app

start:
	docker compose up
stop:
	docker compose down