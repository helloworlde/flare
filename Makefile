binary:
	go run build/build.go
	go build -o bin/flare main.go

test:
	make binary
	./bin/flare

image:
	sh ./build/docker-image-build.sh