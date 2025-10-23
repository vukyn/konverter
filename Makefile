# Makefile for konverter

run:
	go run main.go

build:
	go build -o konverter main.go

docker-build:
	docker build -t konverter .

docker-run:
	docker run -p 6001:8080 -d --name konverter konverter