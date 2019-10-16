FROM golang:1.12

ADD main.go .

RUN go build -o /server .

EXPOSE 80

CMD ["/server"]