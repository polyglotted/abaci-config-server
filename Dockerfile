FROM alpine:3.3
COPY abaci-config-server /server
EXPOSE 8080
ENV RAW_URL=https://github.com

RUN apk add --no-cache ca-certificates

CMD /server -url="$RAW_URL"
