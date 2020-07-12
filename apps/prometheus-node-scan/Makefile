docker-build:
	DOCKER_BUILDKIT=1 docker build -t prometheus-node-scan .

docker-run: docker-build
	-docker rm -f prometheus-node-scan
	docker run -d -p 9090:9090 --name prometheus-node-scan prometheus-node-scan
