FROM golang:1.11.2

RUN mkdir /app

ADD . /app/

WORKDIR /app

ARG SESSION_KEY
ENV SESSION_KEY $SESSION_KEY

RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/gorilla/sessions

RUN go build -o main .

CMD ["/app/main"]
