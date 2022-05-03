FROM debian:buster-slim

COPY .build/app /usr/bin/app

RUN chmod +x /usr/bin/app

ENTRYPOINT ["/usr/bin/app"]
