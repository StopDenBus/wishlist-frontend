build:
	go build -o bin/wishlist-frontend cmd/main.go

run: build
	./bin/wishlist-frontend

.PHONY: build run