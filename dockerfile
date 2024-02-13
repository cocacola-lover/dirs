FROM golang:1.22

WORKDIR /usr/src/dirs-client

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/client

EXPOSE 3333
EXPOSE 3334

CMD ["app"]