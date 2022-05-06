up:
	docker-compose up -d

down:
	docker-compose down

env:
	cp ./.env.example ./.env

migrations:
	go run main.go MakeMigrations

html:
	google-chrome stuff/html/index.html

server:
	go run main.go httpserver

