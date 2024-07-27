FROM golang:alpine AS app

WORKDIR /go/src/roleclub-website/auth/

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main cmd/auth/main.go


FROM alpine

WORKDIR /go/src/roleclub-website/auth/

COPY --from=app /go/src/roleclub-website/auth/ .

EXPOSE 8081

ENTRYPOINT [ "./main" ]