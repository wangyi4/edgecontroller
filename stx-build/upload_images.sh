#! /bin/sh

set -x

REPO=edgecontroller-vm.sh.intel.com:5000
VER=19.06

# tag images
docker tag cce:latest $REPO/edgecontroller-cce:$VER
docker tag ui:latest $REPO/edgecontroller-ui:$VER
docker tag cups:latest $REPO/edgecontroller-cups:$VER

# upload images
docker push $REPO/edgecontroller-cce:$VER
docker push $REPO/edgecontroller-ui:$VER
docker push $REPO/edgecontroller-cups:$VER
