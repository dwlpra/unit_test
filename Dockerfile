FROM golang:alpine

WORKDIR /app 

COPY . .

# Menjalankan unit test
RUN go test -cover ./...

RUN go build -o unit_test 

ENTRYPOINT ["./unit_test"]