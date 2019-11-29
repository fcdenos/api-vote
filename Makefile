#! /bin/bash

build:
	docker build -t ritoon/api-vote:latest .

local-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	docker build -f Dockerfile.local -t ritoon/api-vote:latest .
	make clean-bin
	make run

run:
	#docker run -p 8081:8081 ritoon/api-vote:latest
	docker-compose up

run-local:
	make clean-docker
	docker run -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -p 5432:5432 -d postgres
	sleep 3
	go run main.go

test-vegeta:
	echo "POST http://localhost:8081/v1/user" | vegeta attack -rate 5000 -body ./body.json -header 'Content-Type: application/json' -duration=10s > results.bin
	vegeta report -type=json results.bin > metrics.json
	# echo "POST http://localhost:8081/v1/user" | vegeta attack -rate 50 -body ./body.json -header 'Content-Type: application/json' -duration=1000s | tee results.bin | vegeta report vegeta report -type=json results.bin > metrics.json
	cat results.bin | vegeta plot > plot.html
	cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"

clean-files:
	rm results.bin plot.html metrics.json

clean-bin:
	rm ./app

clean-docker:
	docker stop `docker ps -a -q`
	docker rm `docker ps -a -q`