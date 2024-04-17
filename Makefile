docker-build:
	docker build . -t  ghcr.io/julesrosier/stage-2024-dashboard:latest --build-arg GIT_COMMIT=$$(git log -1 --format=%h)

docker-push:
	docker push ghcr.io/julesrosier/stage-2024-dashboard:latest

docker-update:
	@make --no-print-directory docker-build
	@make --no-print-directory docker-push

codegen:
	templ generate
	sqlc generate

build:
	@make --no-print-directory codegen
	go build -o ./tmp/main.exe ./cmd/server

start:
	@./tmp/main.exe