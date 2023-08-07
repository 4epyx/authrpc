FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go build -o app app.go


FROM alpine

WORKDIR /server
COPY  --from=builder /build/app app
CMD [ "./app" ]
