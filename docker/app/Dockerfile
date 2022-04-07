FROM golang:1.17-alpine

WORKDIR /go/src/app
COPY ./src/api .

RUN go get -u github.com/cosmtrek/air && \
    go get github.com/gin-gonic/gin && \
    go get github.com/jinzhu/gorm && \
    go get github.com/jinzhu/gorm/dialects/postgres && \
    go get github.com/google/uuid

RUN go mod tidy
RUN go mod download

RUN apk update && \
    apk add --no-cache git

# EXPOSE 8080
#RUN echo $PATH
#RUN echo $GOPATH

CMD ["air", "-c", ".air.toml"]