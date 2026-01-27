start-process:
	docker-compose up -d
	go build .
	./login-api


stop-process:
	docker-compose down
