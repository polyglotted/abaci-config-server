FROM scratch
COPY abaci-config-server /server
EXPOSE 8080
VOLUME /data
ENTRYPOINT ["/server"]
