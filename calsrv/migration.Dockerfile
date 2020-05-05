# base the docker container off of the official golang image
FROM golang:latest

# install goose
RUN go get 'bitbucket.org/liamstask/goose/cmd/goose'

# mount the app
RUN mkdir -p /opt/db
ADD ./migration/dbconf.yml /opt/db/dbconf.yml
ADD ./migration/goose /opt/db/migrations
ADD ./wait.sh /opt

# set the working directory to /opt/
WORKDIR /opt/

CMD ["/go/bin/goose", "--env=docker","up"]
