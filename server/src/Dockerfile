FROM golang:1.16.5

RUN mkdir /app
ADD server /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]