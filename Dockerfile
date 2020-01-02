FROM golang:latest

WORKDIR /usr/src/app
COPY . .

RUN go get -d .
RUN go build -o main .

EXPOSE 4500

CMD ./main
