FROM golang:1.15.1

# RUN apt-get -y update
# RUN apt-get -y upgrade
# RUN apt-get install -y sqlite3 libsqlite3-dev build-essential

RUN apt-get update -qq && apt-get install -y default-mysql-client vim

RUN curl -sL https://raw.githubusercontent.com/objectbox/objectbox-go/master/install.sh | bash

WORKDIR /go/src

ADD . /go/src

ENV GO111MODULE=on

EXPOSE 8080

CMD ["go", "run", "/go/src/main.go"]
