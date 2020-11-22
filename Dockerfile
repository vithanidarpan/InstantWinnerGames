# Start from a Debian image with the latest version of Go installed
FROM golang:1.14 as builder

RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get gorm.io/driver/mysql

ADD . /InstantWinnerGames

WORKDIR "/InstantWinnerGames"

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o .

FROM alpine

COPY --from=builder /InstantWinnerGames/InstantWinnerGames /

EXPOSE 8080

CMD ["./main"]