# Docker

Deploy database + API:

1. cd docker
2. docker-compose up

Deploy cron worker:

1. cd docker
2. docker build .. -f Dockerfile_cron -t cron
3. docker run -it --network=docker_default cron