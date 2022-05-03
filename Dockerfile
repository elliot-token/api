FROM gcr.io/distroless/base

COPY --chmod=755 .build/app /usr/bin/app
USER nonroot:nonroot

ENTRYPOINT ["/usr/bin/app"]
