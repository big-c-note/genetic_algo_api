### Dockerize your own image

This assumes you have a Docker Hub account and have docker installed on your machine.

See top level directory README.md for installation instructions for docker.

``` bash
cd path/to/backend_api
docker build -t backend_api .
# Poke around the image to make sure things look good.
docker container run -ti backend_api /bin/sh
# Create new repo on Docker Hub
docker image ls
# Grab the image ID
docker tag {image id} {repo you just made}
# You need sudo on linux.
sudo docker push {repo you just made}
```

### Run backend api locally.

Start by reading consumer_api/README.md, then come back.

``` bash
cd path/to/backend_api
python3 backend_api/api.py
```

Once this api is running, go back to consumer_api/README.md
