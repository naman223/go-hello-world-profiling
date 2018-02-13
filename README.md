# go-hello-world
A tiny hello world go http application that is built into a docker container.

docker build -t go-hello-world:1.0 -f Dockerfile .

docker run -p 9999:9999 go-hello-world:1.0
