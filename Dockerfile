FROM alpine:latest
ENV TZ Asia/Jakarta
EXPOSE 8080
COPY ./bin/ /
COPY ./files/etc/skripsionline-be /
ENTRYPOINT ["/skripsionline-be"]
