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