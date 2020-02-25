to build and put the source file in the docker folder, run
go build -o ./docker/microservice
in the root dir
bazel does not support gomodules, using go compiler for now
to build and run the image locally(not in minikube)
docker build -f 'docker/Dockerfile' -t tmp/microsvc:1.0 .
docker run -p 8080:8080 --rm -d --name micro tmp/microsvc:1.0

k apply -f namespace.yaml
k apply -f resources.yaml

k apply -k overlays/dev
