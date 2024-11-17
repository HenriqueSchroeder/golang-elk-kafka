# Use bash syntax
SHELL := /bin/bash

# Set timezone and user/group ids
export TZ=America/Sao_Paulo
export USER_ID=$(shell id -u)
export GROUP_ID=$(shell if [ `id -g` == 20 ]; then echo '1000'; else echo `id -g`; fi)

# Run all containers in detached mode
run:
	docker-compose up -d

# Stop all containers related to the app
stop:
	docker-compose stop

# Clean: stop and remove containers, networks, volumes, and images
clean:
	docker-compose down -v --remove-orphans

# Show current status of containers
status:
	docker-compose ps

# Restart containers
restart:
	docker-compose restart

# Check logs
logs:
	docker-compose logs -f
