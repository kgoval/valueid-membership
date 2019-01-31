FROM golang:1.11.5-alpine as builder
# Install git + SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates
# Force the go compiler to use modules
ENV GO111MODULE=on
ENV GOPROXY=http://35.240.129.72
# Create appuser
RUN adduser -D -g '' appuser
COPY ./ ./go/app
#COPY ./keys /keys
WORKDIR ./go/app
#RUN go get
RUN go mod download
#get dependancies
#RUN go get
#build the binary
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/apps
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/apps
# STEP 2 build a small image
# start from scratch
FROM alpine
#COPY --from=builder /keys /keys
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
#COPY --from=builder /go/src/golang.myvalue.id/sso/static/ /static/
COPY --from=builder /go/bin/apps /go/bin/apps
USER appuser
#CMD ["tail -f /dev/null"]
#CMD ["./go/bin/apps"]