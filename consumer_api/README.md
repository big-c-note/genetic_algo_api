### Dockerize your own image

This assumes you have a Docker Hub account and have docker installed on your machine.

See top level directory README.md for installation instructions for docker.

``` bash
cd path/to/consumer_api
docker build -t consumer_api .
# Poke around the image to make sure things look good.
docker container run -ti consumer_api /bin/sh
# Create new repo on Docker Hub
docker image ls
# Grab the image ID
docker tag {image id} {repo you just made}
# You need sudo on linux.
sudo docker push {repo you just made}
```

### How to run the api locally.

``` bash
cd path/to/consumer_api
go run .
```


#### Curl Commands.

Below is an example that will work with the genetic algorithm api.

-- Post Request 
``` bash
curl localhost:8080/things -X POST -d '{"name":"laptop","value":500,"weight":2200}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"headphones","value":150,"weight":160}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"coffee mug","value":60,"weight":350}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"notepad","value":40,"weight":333}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"water bottle","value":30,"weight":192}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"mint","value":5,"weight":25}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"socks","value":10,"weight":38}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"tissues","value":15,"weight":80}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"phone","value":500,"weight":200}' -H "Content-Type: application/json"
curl localhost:8080/things -X POST -d '{"name":"baseball cap","value":100,"weight":70}' -H "Content-Type: application/json"
```
 
Once you get things added to the consumer-facing api, run the following get request
to make sure the data looks good.

-- Get Request
``` bash
curl localhost:8080/things
```

Returns JSON of things that have been posted.


Now go checkout ../backend_api/README.md to run the flask api locally. Then come back
to run the algorithm with the below Get Request.

-- Run Algo
``` bash
curl localhost:8080/things/algo
```

This get request sends a get request to the genetic algorithm api.
