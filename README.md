# go-hello-world
A tiny hello world go http application with pprof profiling that is built into a docker container.

docker build -t go-hello-world:1.0 -f Dockerfile .

docker run -p 9999:9999 go-hello-world:1.0

App:
http://localhost:9999

For profiling:
http://localhost:9999/debug/pprof/heap

Now use go tool pprof to check profiling.
