
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
~~~bash task:docker:status
docker stats
~~~

## Git
~~~bash task:git:git-push -- Git add all & commit & push
#args> ct:commit text
git add .
git commit -m "$ct"
git push
~~~

~~~bash task:git:set-git-tag -- set git tag & push to GitHub
#args> t:tag
git tag $t
git tag
git push origin $t
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

~~~bash task:install-guide:linux-amd64 -- Display command (amd64 arch)
#task> _install:make-linux-installer -- arch=amd64
~~~
~~~bash task:install-guide:linux-arm64 -- Display command (arm64 arch)
#task> _install:make-linux-installer -- arch=arm64
~~~

~~~bash task:_install:make-linux-installer
#args> arch=(amd64|arm64)
echo "run command: sudo cp ./sources/mdtk_bin/linux_$arch/mdtk /usr/local/bin/mdtk"
echo 'create script'
filename=install.sh
mdtk _install install-linux --all-task --script -- arch=$arch > $filename
sudo chmod +x $filename
~~~
~~~bash task:_install:install-linux
#args> arch=(amd64|arm64)
echo "Install mdtk (os: linux, arch: $arch)"
sudo cp ./sources/mdtk_bin/linux_$arch/mdtk /usr/local/bin/mdtk
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
