#!/bin/bash
VERSION=$(git describe --tags --abbrev=0)
COMMIT=$(git rev-parse --short HEAD)

echo "最近版本；$VERSION / $COMMIT"

if [[ -z "${DOCKERHUB_REPO}" ]]; then
  DOCKERHUB_REPO="${DOCKERHUB_REPO}"
else
  DOCKERHUB_REPO="soulteary/flare"
fi

docker images | grep "$DOCKERHUB_REPO"

docker manifest rm "$DOCKERHUB_REPO:latest"

docker manifest create --amend "$DOCKERHUB_REPO:latest" \
                               "$DOCKERHUB_REPO:$VERSION-amd64" \
                               "$DOCKERHUB_REPO:$VERSION-arm32v7" \
                               "$DOCKERHUB_REPO:$VERSION-arm64v8" 
docker manifest push "$DOCKERHUB_REPO:latest"
