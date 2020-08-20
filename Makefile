default:
	@echo "=============Building============="
	docker build -f worker/Dockerfile.dev -t worker .

up: default
	@echo "=============Compose Up============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f