FROM alpine:latest
RUN mkdir /app
WORKDIR /app
ADD dist_linux/ /app
CMD ["/app/server"]