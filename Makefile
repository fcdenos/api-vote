build:
	docker build -t ritoon/api-vote:latest .

run:
	docker run -p 8081:8081 ritoon/api-vote:latest