build-cloud-service:
	cd cloud/build/package
	docker build --tag cloud-service .

build-dev:
	sudo docker-compose up -d

stop-dev:
	sudo docker-compose down
	sudo docker-compose kill