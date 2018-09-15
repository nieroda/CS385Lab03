.PHONY: goapi cleandangling

goapi:
	docker build -t goapi ./goapi
	docker run -d --name goapi goapi

cleandangling:
	docker images -qf dangling=true | xargs docker rmi | true

clean:
	docker stop goapi | true
	docker container ls -a -q | xargs docker container rm | true
	docker image ls -a -q | xargs docker image rm | true
