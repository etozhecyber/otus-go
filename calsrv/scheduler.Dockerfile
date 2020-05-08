#build container
FROM golang:alpine AS build-env
ADD . /src
WORKDIR /src
RUN go get
RUN go build -o scheduler apps/scheduler/main.go

#run container
FROM alpine
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=build-env /src/scheduler /app/
COPY --from=build-env /src/config.json /app/
COPY ./wait.sh /app/
CMD ./scheduler
