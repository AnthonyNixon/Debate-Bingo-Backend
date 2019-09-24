FROM golang:1.13-alpine as build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/AnthonyNixon/Debate-Bingo-Backend

COPY go.mod .
COPY go.sum .
RUN go mod tidy

FROM build_base as binary_builder

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /debate-bingo-api .

FROM alpine

COPY --from=binary_builder /debate-bingo-api /debate-bingo-api
ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080

CMD ["./debate-bingo-api"]