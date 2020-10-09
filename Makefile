run: stop up

stop:
	docker-compose -f docker-compose.yml stop

up:
	docker-compose -f docker-compose.yml up -d --build

build:
	docker build -t card-game/listd:1.0 -f cmd/deploy/prod.dockerfile .

test_app:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes


test_continuous:
	docker-compose -f docker-compose.test.yml up --build
	docker-compose -f docker-compose.test.yml down --volumes
