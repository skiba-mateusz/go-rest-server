FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . ./ 
RUN go build -o /bin/go-rest-server ./cmd/main.go
EXPOSE 8080
CMD ["/bin/go-rest-server"]

