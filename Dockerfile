FROM scratch
COPY abaci-config-server /server
EXPOSE 8080
ENV RAW_URL=https://github.com

CMD /server -url="$RAW_URL"
