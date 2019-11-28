build:
	docker build -t ritoon/api-vote:latest .

local-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	docker build -f Dockerfile.local -t ritoon/api-vote:latest .
	rm app
	make run

run:
	#docker run -p 8081:8081 ritoon/api-vote:latest
	docker-compose up