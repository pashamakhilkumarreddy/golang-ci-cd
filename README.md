# Golang CI CD Pipeline

### To run the application in local

- `go run main.go` or
- `go build -o main main.go` and `./main`

### To run the application in local with Docker (make sure Docker is installed)

- `docker build --no-cache -f Dockerfile -t golang-ci-cd .`

- `docker run -dp 8080:8080 golang-ci-cd`

### To push the images to docker hub

- `docker buildx create --name mydockerbuilder --driver docker-container --bootstrap`

- `docker buildx use mydockerbuilder`

- `docker buildx inspect`

- `docker login --username <docker-hub-username> --password <docker-hub-password>`

- `docker buildx build --no-cache --platform=linux/arm64 --platform=linux/amd64 -t docker.io/<docker-hub-username>/<image-name>:<tag> --push -f ./Dockerfile .` or `docker build --no-cache --platform=linux/arm64,linux/amd64 -f Dockerfile --push -t <docker-hub-username>/<image-name> .`
