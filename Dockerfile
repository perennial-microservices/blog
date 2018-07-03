FROM scratch

EXPOSE 6767
ADD accountservice-linux-amd64 /

ADD healthcheck-linux-amd64 /
HEALTHCHECK --interval=30s --timeout=30s --start-period=30s CMD ["./healthcheck-linux-amd64", "-port=6767"] || exit 1

ENTRYPOINT ["./accountservice-linux-amd64"]