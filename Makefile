project = gitlab-dashboard

.PHONY: build
build:
	docker build -t $(project) .

.PHONY: run
run:
	docker run -ti \
		-v $(PWD):/app \
		-e GITLAB_TOKEN=$(token) \
		--network host \
		$(project) \
		go run cmd/main.go

.PHONY: test
test:
	docker run -ti \
		-v $(PWD):/app \
		-e GITLAB_TOKEN=$(token) \
		$(project) \
		go test ./...