.PHONY: up reset clean check-env 

# To run docker compose.
up:
	docker compose up

# To delete both compose containers & volumes.
reset: 	
	docker compose rm && docker compose rm -v

# To delete compose volumes.
clean: 
	docker compose rm -v

# To build
build:
	go build -o dist/mocksy cmd/mocksy/main.go

# To build and run the project
run: 
	go run cmd/mocksy/main.go

