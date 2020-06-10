FROM alpine:latest

COPY hello /root

ENTRYPOINT ["/root/hello"]
