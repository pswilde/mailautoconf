FROM golang:1-alpine3.14 AS builder

COPY src/ /mailautoconf
WORKDIR /mailautoconf
RUN go build -o /mailautoconf/mailautoconf

FROM alpine:3.14

RUN apk add --no-cache bash
COPY --from=builder /mailautoconf/mailautoconf /mailautoconf/mailautoconf
COPY --from=builder /mailautoconf/default-config /mailautoconf/default-config
COPY --from=builder /mailautoconf/templates /mailautoconf/templates

COPY ./entrypoint.sh /
RUN chmod +x /entrypoint.sh

EXPOSE 8010

ENTRYPOINT ["/entrypoint.sh"]
