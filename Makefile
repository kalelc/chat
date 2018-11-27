REPOSITORY = kalelc/chat
NAME = chat

build:
	docker build --build-arg SESSION_KEY=$(SESSION_KEY) -t $(NAME) -f Dockerfile .
run:
	docker run -p 8000:8000 $(NAME)
tag: build
	docker tag $(NAME) $(REPOSITORY)
push: tag
	docker push $(REPOSITORY)
