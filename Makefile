build:
	docker build --tag chelout/otus-homework_1:latest .

push:
	docker push chelout/otus-homework_1:latest

run:
	docker run -p 8000:8000 chelout/otus-homework_1:latest
