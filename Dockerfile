FROM golang:alpine AS builder

# Stage I
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN go build -o /go/bin/gensectext

# Stage II
FROM scratch
COPY --from=builder /go/bin/gensectext /go/bin/gensectext

ENTRYPOINT ["/go/bin/gensectext"]
