FROM golang:1.19
EXPOSE 8080
WORKDIR /cmd
COPY . .
RUN go build -o main .
CMD ["/cmd/main"]