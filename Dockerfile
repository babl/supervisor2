FROM busybox

EXPOSE 4445
ADD supervisor2_linux_amd64 /bin/supervisor2
RUN chmod +x /bin/supervisor2

CMD ["/bin/supervisor2", "--kafka-brokers=queue.babl.sh:9092"]
