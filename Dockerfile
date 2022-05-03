FROM gcr.io/distroless/base

COPY .build/app /usr/bin/app

USER nonroot:nonroot

ENTRYPOINT ["/usr/bin/app"]
