FROM alpine:edge

RUN apk add --no-cache sassc

COPY start.sh /start.sh

VOLUME ["/src", "/dst"]

CMD ["/start.sh"]