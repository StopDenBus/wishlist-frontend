FROM alpine:3.22

RUN addgroup -S app -g 1001 && adduser -S app -G app -u 1001

USER app

WORKDIR /app

COPY --chown=app:app . /app/

RUN chmod u+x /app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]