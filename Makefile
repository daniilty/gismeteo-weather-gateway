build:
	go build -o server github.com/daniilty/gismeteo-weather-gateway/cmd/server
build_docker:
	docker build -t gismeteo:latest -f docker/Dockerfile .
