.PHONY: build test run docker-up docker-down clean

watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching for file changes...";\
	else \
	  read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
      	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
      	        go install github.com/air-verse/air@latest; \
      	        air; \
      	        echo "Watching...";\
      	    else \
      	        echo "You chose not to install air. Exiting..."; \
      	        exit 1; \
      	    fi; \
    fi

build:
	@echo "Building..."
	@go build -o main cmd/app/main.go

test:
	@echo "Testing..."
	@go test ./tests -v

run:
	@echo "Starting server..."
	@go run cmd/app/main.go

#Boot up all containers
docker-up:
	@if docker compose up --build 2>/dev/null; then \
  		: ; \
  	else \
  	  echo "Using docker compose version 1"; \
  	  docker-compose up; \
  	fi

#Stop containers
docker-down:
	@if docker compose down 2>/dev/null; then \
      		: ; \
      	else \
      	  echo "Using docker compose version 1"; \
      	  docker-compose down; \
      	fi

clean:
	@echo "Cleaning out the closet...🤫"
	@echo "This will remove the binary file and all containers + volumes created by this system..."
	@rm -f main
	@if docker compose down -v 2>/dev/null; then \
          	: ; \
		else \
        	echo "Using docker compose version 1"; \
          	  docker-compose down -v; \
          	fi