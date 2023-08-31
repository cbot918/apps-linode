NAME=getsub

drun: dbuild
	docker run -dit -p 80:8080 --name $(NAME) $(NAME)
	
dbuild: dockerfile
	docker build -t $(NAME) .

drmall:
	drmc $(NAME)
	drmi $(NAME)