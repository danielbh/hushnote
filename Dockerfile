FROM golang:alpine as builder
RUN apk add --no-cache git
COPY . /src
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo
# RUN go build

FROM scratch
COPY --from=builder ./src/go-bank ./app
CMD ["./app"]
