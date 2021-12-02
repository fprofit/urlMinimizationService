# urlMinimizationService
go run ./cmd/main.go -inMemory 
go run ./cmd/main.go -dbSQL

# docker
docker-compose up

docker run -p 8080:8080 urlminimizationservice_app //start servise inMemory
docker-compose start //start servise dbSQL