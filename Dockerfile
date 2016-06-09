FROM scratch
COPY server /
EXPOSE 8080
VOLUME /data
ENTRYPOINT ["/server"]
