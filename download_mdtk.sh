#!/bin/bash
set -euo pipefail

arch=$1

if [ $arch != "amd64" ] && [ $arch != "arm64" ]; then
    echo arch \('$1'\) is amd64 or arm64.
    exit 1
fi 

j=$(wget -qO - https://api.github.com/repos/m-YoC/mdtk/releases/latest)
fname=$(echo $j | jq -r ".assets[] | select(.name | test(\"^mdtk_bin_v.+_$arch.tar.gz$\")) | .name")
url=$(echo $j | jq -r ".assets[] | select(.name | test(\"^mdtk_bin_v.+_$arch.tar.gz$\")) | .browser_download_url")

(
    cd $(dirname $0)

    wget -q $url
    tar -zxvf ./$fname
    rm -R ./$fname
)
