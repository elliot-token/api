FROM gcr.io/distroless/base

COPY /.build/app /usr/bin/app

ENTRYPOINT ["/usr/bin/app"]
