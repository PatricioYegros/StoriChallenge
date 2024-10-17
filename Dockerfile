FROM golang
WORKDIR /path

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

#RUN CGO_ENABLED=0 GOOS=windows go build -o /storichallenge/app
#EXPOSE 8080
#CMD ["/storichallenge/app"]

FROM alpine:3.19
COPY --from=builder /app/stori_challenge stori_challenge
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot
USER nonroot
ENTRYPOINT ["/stori_challenge"]