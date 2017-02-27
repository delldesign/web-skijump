FROM alpine:latest

RUN apk add --no-cache --virtual .build-deps curl tar \
      && curl -L https://github.com/spf13/hugo/releases/download/v0.18.1/hugo_0.18.1_Linux-64bit.tar.gz | tar -xz \
      && mv /hugo_0.18.1_linux_amd64/hugo_0.18.1_linux_amd64 /usr/bin/hugo \
      && rm -r /hugo_0.18.1_linux_amd64/ \
      && apk del --no-cache .build-deps

EXPOSE 1313

VOLUME /src

ENTRYPOINT ["/usr/bin/hugo", "-s", "/src", "server", "-w", "--bind", "0.0.0.0"]