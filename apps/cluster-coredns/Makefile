.PHONY: docker-build-local
docker-build-local:
	@DOCKER_BUILDKIT=1 docker build -t cluster-coredns .

.PHONY: docker-run-local
docker-run-local: docker-build-local
	@docker rm -f cluster-coredns 2> /dev/null || true
	@docker run -d -p 53:53 -p 53:53/udp -p 8080:8080 -p 9153:9153 --name cluster-coredns cluster-coredns

.PHONY: docker-build-and-push
docker-build-and-push:
	@docker buildx build --platform linux/amd64,linux/arm/v7 -t rafaeleyng/cluster-coredns:latest --push .

.PHONY: docker-run
docker-run:
	@docker rm -f cluster-coredns 2> /dev/null || true
	@docker run -d -p 53:53 -p 53:53/udp -p 8080:8080 -p 9153:9153 --name cluster-coredns --restart=unless-stopped rafaeleyng/cluster-coredns
