FROM golang:1.17.3-buster

EXPOSE 8080:8080

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o minimizatorUrl ./cmd/main.go

CMD ["./minimizatorUrl", "-inMemory"]

