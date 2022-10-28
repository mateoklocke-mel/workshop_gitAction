FROM golang:1.16-alpine

WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY ./main.go ./

RUN go build -o /docker-workshop-action

CMD [ "/docker-workshop-action" ]