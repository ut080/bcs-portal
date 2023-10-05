FROM golang:1.21.1

WORKDIR /usr/src/bcs-portal

COPY ../../../go.mod ../../../go.sum ./
RUN go mod download && go mod verify

COPY ../../../ ./
