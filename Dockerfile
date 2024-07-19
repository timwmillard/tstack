FROM golang:alpine

WORKDIR $GOPATH/app


COPY go.mod go.sum ./
RUN go mod download

# COPY *.go *.html *.css *.js ./
COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o /appd ./cmd/server

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

ENTRYPOINT ["/appd", "-migrate"]


