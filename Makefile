build-docker:
	@docker build -t portfoligo .

run-docker:
	@docker container run -d --name portfoligo -p 8080:8080 portfoligo

all: clean build-docker run-docker

clean:
	@docker container stop portfoligo
	@docker container remove portfoligo