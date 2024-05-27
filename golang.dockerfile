
FROM golang:1.22.1-bookworm

RUN apt-get update && apt-get install -y --no-install-recommends  make && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ARG WORKDIR="mdtk"
RUN mkdir /$WORKDIR && umask 0000

WORKDIR /$WORKDIR

ENV TZ=Asia/Tokyo

ENTRYPOINT ["/bin/bash"]

