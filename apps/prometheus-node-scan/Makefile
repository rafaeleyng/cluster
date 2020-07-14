.PHONY: docker-build-local
docker-build-local:
	@DOCKER_BUILDKIT=1 docker build -t prometheus-node-scan .

.PHONY: docker-run-local
docker-run-local: docker-build-local
	@docker rm -f prometheus-node-scan 2> /dev/null || true
	@docker run -d -p 9090:9090 --name prometheus-node-scan prometheus-node-scan

.PHONY: docker-build-and-push
docker-build-and-push:
	@docker buildx build --platform linux/amd64,linux/arm/v7 -t rafaeleyng/prometheus-node-scan:latest --push .

.PHONY: docker-run
docker-run:
	@docker rm -f prometheus-node-scan 2> /dev/null || true
	@docker run -d -p 9090:9090 --name prometheus-node-scan --restart=unless-stopped rafaeleyng/prometheus-node-scan
