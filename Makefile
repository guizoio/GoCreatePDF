up:
	docker-compose up -d

down:
	docker-compose down

env:
	cp ./.env.example ./.env

migrations:
	go run main.go MakeMigrations

server:
	go run main.go httpserver

