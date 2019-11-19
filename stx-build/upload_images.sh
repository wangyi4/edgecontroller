#! /bin/sh

set -x

REPO=edgecontroller-vm.sh.intel.com:5000
VER=19.06

# tag images
docker tag cce:latest $REPO/cce:$VER
docker tag ui:latest $REPO/ui:$VER
docker tag cups:latest $REPO/cups:$VER

# upload images
docker push $REPO/cce:$VER
docker push $REPO/ui:$VER
docker push $REPO/cups:$VER
