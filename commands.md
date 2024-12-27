# Docker

- `docker build --tag bande-a-part .`
    - build the docker image
- `docker build -t bande-a-part-api:multistage -f Dockerfile.multistage .`
    - build a multistage image
- `docker image ls`
    - list all images
- `docker image rm bande-a-part`
    - remove the image

- `docker run image`
    - run the docker image
    - `--publish <machine port>:<docker port>` (`-p`)
        - connect the exposed port to the corresponding local port
    - `--detach` (`-d`)
        - detach the run of the container from the terminal

- `docker ps -a`
    - list all containers

- `docker rm $(docker ps -qa)`
    - delete all non running containers

- `docker container start <container name>`
    - start an already build container
    - replace `start` with `stop` to stop it

# MongoDB

## deprecated, use the docker instead
- fichier de config à `/etc/mongod.conf`
- start mongodb process:
    - `sudo systemctl start mongod`
    - if fails to start:
        - `sudo systemctl daemon-reload`
- check if start successfully:
    - `sudo systemctl status mongod`

## docker version

```sh
docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server
# any data created during the run of the container will be destroyed if the container is destroyed
docker stop mongodb && docker rm mongodb
# to persist the data on the system
docker run --name mongodb -d -p 27017:27017 -v bande-a-part-db:/data/db mongodb/mongodb-community-server

```

### Docker volumes 

Docker volumes are not necessarily on the path you give, you just give a name and it will be a space handled by docker
