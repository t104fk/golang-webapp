FROM golang:1.4.2
MAINTAINER takasing
COPY . /go/src/golang-webapp
WORKDIR /go/src/golang-webapp

ENV ENVIRONMENT=development \
  DB_HOST=db \
  DB_PORT=5432 \
  DB_NAME=toyo \
  DB_USER=takasing \
  DB_PASS=oifwslu4SGcBiMzL

RUN make dependency && make install
EXPOSE 3000
CMD ["golang-webapp"]
