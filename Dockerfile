FROM alpine:3.3

RUN apk add --no-cache ca-certificates
COPY abaci-config-server /server

EXPOSE 8080

ENV PORT=8080 RAW_URL=https://github.com LOCAL_VOL=false

CMD /server -port=$PORT -url="$RAW_URL" -local=$LOCAL_VOL
