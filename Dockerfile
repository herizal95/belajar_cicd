FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /belajar_cicd

EXPOSE 8080

CMD ["/belajar_cicd"]
