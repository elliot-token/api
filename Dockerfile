FROM gcr.io/distroless/base

ARG GIT_COMMIT

COPY /build/app /usr/bin/app

ENTRYPOINT ["/usr/bin/app"]