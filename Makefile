all:	depensure build

build:
	go build -o bin/s3proxy

depensure:
	dep ensure

s3up:	export SERVICES=s3
s3up:
	docker-compose up -d localstack

s3down:
	docker-compose rm -f -s localstack

clean:	s3down
	@rm -rf s3proxy