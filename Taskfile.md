
## Docker

~~~bash task:docker:up -- Start all container
docker compose up -d
~~~
~~~bash task:docker:down -- Down all container
docker compose down
~~~
~~~bash task:docker:build -- Build all container
docker compose build
~~~
~~~bash task:docker:build-plain -- Build all container (--progress=plain)
docker compose build --progress=plain
~~~

## Compress to .tar.gz and Decompress

~~~bash task:tar.gz:compress -- Compress binary files
cd sources
source ./mdtk/version.txt
tar -zcvf ../mdtk_bin_v${VERSION}.tar.gz ./mdtk_bin 
~~~
~~~bash task:tar.gz:decompress -- Decompress binary files
echo 'run command: tar -zxvf ./mdtk_bin_VERSION.tar.gz'
~~~

## Install Guide
[!] Cannot install mdtk using mdtk.

~~~bash task:install-guide:linux-amd64 -- Display command (amd64 arch)
echo 'run command: sudo cp ./sources/mdtk_bin/linux_amd64/mdtk /usr/local/bin/mdtk'
~~~
~~~bash task:install-guide:linux-arm64 -- Display command (arm64 arch)
echo 'run command: sudo cp ./sources/mdtk_bin/linux_arm64/mdtk /usr/local/bin/mdtk'
~~~

## Utils

If the task content is not clear and the description is not written outside the code block,   
it is better to use the `#desc>` notation for the task description.  
This makes it easier to understand when displaying the preview.

~~~bash task:utils:size 
#desc> Get mdtk binary size (linux amd64 arch)
cd ./sources/mdtk_bin/linux_amd64
ls -lh
~~~
