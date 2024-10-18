FROM golang:alpine AS builder
RUN apk --no-cache add build-base
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o storichallenge
FROM alpine:3.19
COPY --from=builder app/storichallenge storichallenge
#RUN addgroup -g 1000 appuser \
#    && adduser -u 1000 -g 1000 appuser 
#USER appuser
RUN addgroup -S nonroot \
    && adduser -S nonroot -g nonroot
USER nonroot
ENTRYPOINT ["/storichallenge"]