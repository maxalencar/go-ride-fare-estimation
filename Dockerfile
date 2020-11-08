FROM golang:alpine
    
WORKDIR /go/src/go-ride-fare-estimation
COPY . .

RUN cd ./cmd/go-ride-fare-estimation && go build -v ./...

CMD ./cmd/go-ride-fare-estimation/go-ride-fare-estimation