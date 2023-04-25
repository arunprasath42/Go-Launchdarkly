FROM golang:1.16-alpine

WORKDIR /LaunchDarkly

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o LaunchDarkly .

CMD ["./LaunchDarkly"]