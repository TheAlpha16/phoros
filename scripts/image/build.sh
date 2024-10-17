#!/bin/bash

set +x

resource="phoros"
registry="docker.io/thealpha16"

docker images | grep ${registry}/${resource}
echo ""

echo "Choose version to tag"
echo -n "> "
read version
echo ""

echo "tag this to latest version?"
echo "1. YES"
echo "2. NO"
echo -n "> "
read tag
echo ""

docker buildx build --tag ${registry}/${resource}:${version} --platform linux/arm64/v8,linux/amd64 --push .

case $tag in
    "1")
        docker rmi -f ${registry}/${resource}:latest
        docker buildx build --tag ${registry}/${resource}:latest --platform linux/arm64/v8,linux/amd64 --push .
    ;;
    "2")
        echo "[*] Not removing latest tag"
    ;;
    *)
        echo "[-] Choose 'yes' or 'no' for tagging"
        echo ""
        echo "[-] exiting"
        exit
    ;;
esac