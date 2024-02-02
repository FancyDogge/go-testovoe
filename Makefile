build:
	@go build -o bin/api 

run: build
	@./bin/api

docker:
	echo "building docker file"
	@docker build -t api .
	echo "running API inside Docker container"
	@docker run -p 8000:8000 api