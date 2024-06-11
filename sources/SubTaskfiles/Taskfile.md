

```bash task:sub:subtest
#desc> hello mdtk
echo hello mdtk! wd:`pwd`
```

### Make Sample.dockerfile

```bash task::dftest
mdtk dockerfile dockerfile-merge -a > Sample.dockerfile
```


## dockerfiles

### Merge dockerfiles

```dockerfile task:dockerfile:dockerfile-merge
#
ARG TIMEZONE=Asia/Tokyo

# ------------------------------------------------------
#embed> dockerfile:docker
# ------------------------------------------------------
#embed> dockerfile:terraform
# ------------------------------------------------------
#embed> dockerfile:awscliv2
# ------------------------------------------------------

FROM ubuntu:22.04 AS aws-tf
ARG TARGETARCH
ARG TIMEZONE

#embed> dockerfile:docker-install

#embed> dockerfile:terraform-install

#embed> dockerfile:awscliv2-install

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE \ 
    apt-get install -y \ 
    make gawk jq curl tzdata && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Add New Layers...

ENV TZ=$TIMEZONE
```

### AWS CLI v2

```dockerfile task:dockerfile:awscliv2
# Need Args: TIMEZONE
FROM ubuntu:22.04 AS aws-cli-v2
ARG TARGETARCH
ARG TIMEZONE

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE \ 
    apt-get install -y curl unzip
RUN [ "$TARGETARCH" = "amd64" ] && curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"  -o "awscliv2.zip" || true
RUN [ "$TARGETARCH" = "arm64" ] && curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" || true
RUN unzip awscliv2.zip
RUN ./aws/install -i /usr/local/aws-cli -b /usr/local/bin

# COPY --from=aws-cli-v2 /usr/local/aws-cli /usr/local/aws-cli
# COPY --from=aws-cli-v2 /usr/local/bin /usr/local/bin
# COPY ./aws-assume-role.sh /usr/local/bin/aws-assume
```

#### Install
```dockerfile task:dockerfile:awscliv2-install
COPY --from=aws-cli-v2 /usr/local/aws-cli /usr/local/aws-cli
COPY --from=aws-cli-v2 /usr/local/bin /usr/local/bin
```

### Terraform

```dockerfile task:dockerfile:terraform
# Need Args: TIMEZONE
ARG TF_KEYRING_PATH="/etc/apt/keyrings/hashicorp.gpg"
ARG TF_APT_LIST_PATH="/etc/apt/sources.list.d/hashicorp.list"

# ubuntu ver: 14.04, 16.04, 18.04, 20.04, 21.10, 22.04
FROM ubuntu:22.04 AS terraform
ARG TIMEZONE
ARG TF_KEYRING_PATH
ARG TF_APT_LIST_PATH
ARG TF_URI="https://apt.releases.hashicorp.com"

# 環境のインストール
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE \
    apt-get install -y \ 
    wget curl gnupg software-properties-common && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install the HashiCorp GPG key.
RUN wget -O- $TF_URI/gpg | gpg --dearmor -o $TF_KEYRING_PATH

# Verify the key's fingerprint.
RUN gpg --no-default-keyring --keyring $TF_KEYRING_PATH --fingerprint

# Add the official HashiCorp repository to your system. 
# The lsb_release -cs command finds the distribution release codename 
# for your current system, such as buster, groovy, or sid.
RUN echo "deb [signed-by=$TF_KEYRING_PATH] $TF_URI $(lsb_release -cs) main" | \
    tee $TF_APT_LIST_PATH

# Install
# apt-get update &&apt-get install terraform
```

#### Install
```dockerfile task:dockerfile:terraform-install
ARG TF_KEYRING_PATH
ARG TF_APT_LIST_PATH
COPY --from=terraform   $TF_KEYRING_PATH        $TF_KEYRING_PATH
COPY --from=terraform   $TF_APT_LIST_PATH       $TF_APT_LIST_PATH

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE apt-get install -y ca-certificates && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE apt-get install -y \ 
    terraform && \ 
    apt-get clean && rm -rf /var/lib/apt/lists/*
```

### Docker in Docker

```dockerfile task:dockerfile:docker
# Need Args: TIMEZONE
ARG DOCKER_KEYRING_PATH="/etc/apt/keyrings/docker.asc"
ARG DOCKER_APT_LIST_PATH="/etc/apt/sources.list.d/docker.list"

# ubuntu ver: 14.04, 16.04, 18.04, 20.04, 21.10, 22.04
FROM ubuntu:22.04 AS docker
ARG TIMEZONE
ARG DOCKER_KEYRING_PATH
ARG DOCKER_APT_LIST_PATH
ARG DOCKER_URI="https://download.docker.com/linux/ubuntu"

# 環境のインストール
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE \
    apt-get install -y \ 
    curl ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Add Docker's official GPG key:
RUN curl -fsSL $DOCKER_URI/gpg -o $DOCKER_KEYRING_PATH

# Add the repository to Apt sources:
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=$DOCKER_KEYRING_PATH] \ 
    $DOCKER_URI $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
    tee $DOCKER_APT_LIST_PATH

# Install
# apt-get update && apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# In compose.yml
# volumes:
#      - "/var/run/docker.sock:/var/run/docker.sock"
```

#### Install
```dockerfile task:dockerfile:docker-install
ARG DOCKER_KEYRING_PATH
ARG DOCKER_APT_LIST_PATH
COPY --from=docker      $DOCKER_KEYRING_PATH    $DOCKER_KEYRING_PATH
COPY --from=docker      $DOCKER_APT_LIST_PATH   $DOCKER_APT_LIST_PATH

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE apt-get install -y ca-certificates && \
    DEBIAN_FRONTEND=noninteractive TZ=$TIMEZONE apt-get install -y \ 
    docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
```


