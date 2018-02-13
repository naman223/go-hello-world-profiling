# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ADD main.go /home

WORKDIR /home

CMD ["go","run","main.go"]

# Document that the service listens on port 9999.
EXPOSE 9999
