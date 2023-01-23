FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /belajar-cicd

EXPOSE 5000

CMD ["/belajar-cicd"]