FROM golang:1.18.0-alpine3.14

WORKDIR /LaunchDarkly

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .
RUN apk add --no-cache git
RUN go build -o LaunchDarkly .

CMD ["./LaunchDarkly"]