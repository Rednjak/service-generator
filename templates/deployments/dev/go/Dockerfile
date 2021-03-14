FROM golang:1.16.0-alpine as builder

WORKDIR /server

ENV GO111MODULE=on
ENV GOBIN=/server/bin
ENV GOFLAGS=-mod=vendor

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# Create a user without root privileges
RUN adduser -D -g '' server_user

RUN go install ./cmd/server

####------------------------------------------------------####

FROM golang:1.16.0-alpine

# Copy the dependencies from step 1
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /server

COPY --from=builder /server/bin/server .

RUN apk add --no-cache bash

# Use an unprivileged user.
USER server_user

CMD ["./server"]
