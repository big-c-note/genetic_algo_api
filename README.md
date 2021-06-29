# Genetic Algorithm API


## Introduction and Project Description

The following will be a "tutorial" (read: set of instructions) for
how I set up the project in this repository.

The project aims to deploy two containerized API applications into a kubernetes
cluster. One API is written in Go, and it is the "consumer-facing" endpoint. The
other is a "backend" flask API (python) for manipulating a genetic algorithm.

The consumer can post "things" which have associated "weight" and "value". Once
complete, the consumer can request an optimization such that the value is maximized
under a certain weight threshold. A use case would be if you are going on a trip,
and you can only bring X lbs of stuff, but you want to bring the stuff that makes
you happiest (ie; yields maximum "value"). This is the classic knapsack problem.


### Resources Used and Special Thanks.

The following code bases were used to create the genetic alorithm and go api.

- https://github.com/kiecodes/genetic-algorithms
- https://github.com/viveksinghggits/library
- https://github.com/kubucation/go-rollercoaster-api

### Technologies Used

For me, this project is largely practice in creating containerized applications
and deploying them to a kubernetes cluster.

- go rest api
- python flask api
- genetic algorithm written in python
- docker for containerizing the above applications
- kubernetes for deploying the containerized applications (using minikube)
- helm for packaging the deployment to kubernetes

I am on a linux, ubuntu OS (Focal LTS).

### Comsumer-facing API MVP Checklist

The following will be written in go

- [X] Post request for single thing which contains "name",  "value", and "weight".
- [X] Get request for the list of things that have been posted.
- [X] Get request that, when received, runs a separate get request to the "backend" API
to run the genetic algorithm and returns the results. 

### Backend API MVP Checklist

- [X] powers a get request to the consumer-facing api to receive data,
then calls the genetic algorithm and sends a post request to the consumer-facing
api with the results.

### Kubernetes Deployment

- [X] Create docker images.
- [X] Create helm chart.

## Deploy Locally

If you're just looking to run the project locally, without kubernetes, look at
consumer_api/README.md and backend_api/README.md.

## Deploy in Kubernetes

If you want to deploy in kubernetes.

Start by dockerizing the applications in consumer_api and backend_api directories.

There is documentation of this process in each directory's README.

Now start up minikube (installation instructions below).

```
minikube start
```

You can either apply the supplied yaml configs, or install the helm chart.

### Apply the supplied YAML files.

The manifest files are in the manifests/consumer_api directory. After starting 
minikube, they are applied as such.


*Important note: I did not leave up my docker images. If you want to run this
project, you must dockerize the applications and push to your own repo. You'll
need to change out the "image" references in manifests/consumer_api/consumerapi.yaml
to your image repositories.*

``` bash
cd path/to/top/directory
kubectl apply -f manifests/consumer_api/ns.yaml
kubectl apply -f manifests/consumer_api/consumerapi.yaml
kubectl apply -f manifests/consumer_api/service.yaml
```

Finally you can forward your localhost:8080 to the consumer api service to use.

``` bash
kubectl port-forward -n consumer-api svc/consumerapi-service 8080
```
(see play with the api below)

### Run the Helm Chart Installation.

I also provided a helm chart. Install with

``` bash
helm install genalgo helm
```

Finally you can forward your localhost:8080 to the consumer api service to use.

``` bash
kubectl port-forward -n consumer-api svc/consumerapi-service 8080
```
(see play with the api below)


### Finally, play with the api!

I left a set up curl commands in consumer_api/README.md that will show you how to
interact with the api.

Feel free to mess around with the hardcoded parameters set in
backend_api/backend_api/gen_algo.py.


### Tear Down

Just in case you want to clean up after you install the project.

``` bash
minikube stop
minikube delete
docker image ls 
# find an image you want to delete
docker image rm -f {image id}
# you might additionally want to delete the docker repositories you made
```


## Instructions for installing the software needed 

I installed the following for Ubuntu, but there may be instructions included for
another OS in the below links.

#### Docker.

* [Installation instructions](https://docs.docker.com/engine/install/ubuntu/)
    * I used the install from repo instructions.
* [Quickstart](https://docs.docker.com/docker-hub/)
    * Remember to use sudo in front of docker command and that when you login/push
    * the username must be lowercase. This may be linux only. Not sure.

#### Kubectl.

* [Installation instructions](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
    * I used the curl instructions. 

#### MiniKube.

* [Installation instructions](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-minikube/)

##### Troublshooting,

--issue was I needed to do the post install for Docker:
https://docs.docker.com/engine/install/linux-postinstall/


#### Install Helm

* [Installation instructions](https://helm.sh/docs/intro/install/)
    * used the from script instructions.
