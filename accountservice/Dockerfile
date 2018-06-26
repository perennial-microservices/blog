FROM iron/base

EXPOSE 6767
ADD accountservice-linux-amd64 /
RUN chmod +x ./accountservice-linux-amd64
ENTRYPOINT ["./accountservice-linux-amd64"]