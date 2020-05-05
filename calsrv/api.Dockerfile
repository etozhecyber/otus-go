#build container
FROM golang:alpine AS build-env
ADD . /src
WORKDIR /src
RUN go get
RUN go build -o api_server apps/api/main.go

#run container
FROM alpine
WORKDIR /app
COPY --from=build-env /src/api_server /app/
COPY --from=build-env /src/config.json /app/
CMD ./api_server
