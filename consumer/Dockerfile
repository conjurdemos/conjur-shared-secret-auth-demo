# Build golang binary stage
FROM golang:1.10

RUN go get github.com/cyberark/conjur-api-go/conjurapi
COPY ./consumer.go .

RUN touch /etc/conjur.conf

RUN CGO_ENABLED=0 go build -v -o consumer .

# Final container stage
FROM scratch

COPY --from=0 /go/consumer /consumer

ENV CONJUR_VERSION=5

CMD [ "/consumer" ]