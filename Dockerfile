FROM golang:1.15.1

# RUN mkdir /go/src

# RUN apt-get -y update
# RUN apt-get -y upgrade
# RUN apt-get install -y sqlite3 libsqlite3-dev build-essential

WORKDIR /go/src

ADD . /go/src

ENV GO111MODULE=on

# RUN go mod download

EXPOSE 8080

CMD ["go", "run", "/go/src/main.go"]
