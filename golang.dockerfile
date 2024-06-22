FROM golang:1.22.1-bookworm AS pwsh

# https://learn.microsoft.com/ja-jp/powershell/scripting/install/install-ubuntu?view=powershell-7.4

ARG OS_VERSION_ID=22.04
# Install pre-requisite packages.
RUN apt-get update && apt-get install -y wget apt-transport-https software-properties-common
# Download the Microsoft repository keys
RUN wget -q https://packages.microsoft.com/config/ubuntu/$OS_VERSION_ID/packages-microsoft-prod.deb


# -------------------------------------------------------------------

FROM golang:1.22.1-bookworm

COPY --from=pwsh /go/packages-microsoft-prod.deb /go
# Register the Microsoft repository keys. Delete the Microsoft repository keys file
RUN dpkg -i packages-microsoft-prod.deb && rm packages-microsoft-prod.deb

RUN apt-get update && apt-get install -y --no-install-recommends make powershell && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ARG WORKDIR="mdtk"
RUN mkdir /$WORKDIR && umask 0000

WORKDIR /$WORKDIR

ENV TZ=Asia/Tokyo

ENTRYPOINT ["/bin/bash"]

