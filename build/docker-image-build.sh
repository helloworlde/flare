#!/bin/bash
VERSION=$(git describe --tags --abbrev=0)
COMMIT=$(git rev-parse --short HEAD)

echo "最近版本；$VERSION / $COMMIT"
echo "$VERSION">RELEASE_VERSION
echo "$COMMIT">RELEASE_COMMIT

if [[ -z "${DOCKERHUB_REPO}" ]]; then
  DOCKERHUB_REPO="${DOCKERHUB_REPO}"
else
  DOCKERHUB_REPO="soulteary/flare"
fi

docker build -t "flare-base:$VERSION" -f docker/Dockerfile.base .
docker build -t "$DOCKERHUB_REPO:$VERSION-amd64" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/Dockerfile.amd64 .
docker build -t "$DOCKERHUB_REPO:$VERSION-arm32v7" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/Dockerfile.arm32v7 .
docker build -t "$DOCKERHUB_REPO:$VERSION-arm64v8" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/Dockerfile.arm64v8 .

rm RELEASE_VERSION
rm RELEASE_COMMIT

docker images | grep "$DOCKERHUB_REPO"