.PHONY: \
	docker-run

###################
# common
###################
clean:
	@rm -f ./my-remote

run: clean
	test ${SLACK_TOKEN}
	@go run main.go

setup:
	go get github.com/oxequa/realize

###################
# amd64
###################
build: clean
	@go build

watch:
	test ${SLACK_TOKEN}
	realize start --no-config --run

###################
# arm32v6
###################
build-arm: clean
	@env GOOS=linux GOARCH=arm GOARM=6 go build

docker-run-arm:
	test ${SLACK_TOKEN}
	-docker rm -f my-remote
	-docker rmi -f my-remote
	docker build -t my-remote --build-arg slack_token=${SLACK_TOKEN} .
	docker run \
		--log-driver json-file \
		--name my-remote \
		--restart unless-stopped \
		-d \
		-p 9000:9000 \
		my-remote
