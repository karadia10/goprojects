FROM golang:1.20.2-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD [ "/app/main" ]