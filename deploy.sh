go build hello.go
docker build . -t $DOCKER_USERNAME/go-hello
# docker push $DOCKER_USERNAME/go-hello
