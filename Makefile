up:
	docker-compose up -d

down:
	docker-compose down

env:
	cp ./.env.example ./.env

migrations:
	go run main.go MakeMigrations

topic:
	docker exec kafka kafka-topics --topic create.client.v0 --create --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

html:
	google-chrome stuff/html/index.html

server:
	go run main.go httpserver

