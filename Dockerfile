FROM alpine:3.3

RUN apk add --no-cache ca-certificates
COPY abaci-config-server /server

EXPOSE 8080
VOLUME /data

ENV RAW_URL=https://github.com LOCAL_VOL=false

CMD /server -url="$RAW_URL" -local=$LOCAL_VOL
