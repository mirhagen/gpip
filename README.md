# gpip

![Test](https://github.com/RedeployAB/gpip/workflows/Test/badge.svg)

> HTTP service that returns callers IP address

`gpip` is a service to be used for returning callers IP address.
A single endpoint to handle requests.

It supports supports `application/json` and `text/plain` for handling different response types.
To determine the calling IP the headers `Forwarded`, then `X-Forwarded-For` and finally `X-Real-IP` are checked
before falling back to incoming remote address.

**Contents**

* [Endpoints](#endpoints)
* [Configuration](#configuration)
* [Build](#build)
* [Deployment](#deployment)
  * [Kubernetes](#kubernetes)
* [Additional credits](#additional-credits)


## Endpoints

```
// Get IP as JSON
// With Accept: application/json, Accept: */* or no Accept header provided
// it will respond with application/json.
Accept: application/json
GET /

// Response
{"ip":"<ip-address>"}

// Get IP as text/plain
Accept: text/plain
GET /

// Response
<ip-address>
```

## Configuration

`gpip` listens by default on address `0.0.0.0:5050`. To configure host and port, flags
or environment variables can be used.

**Environment variables**

```bash
GPIP_LISTEN_HOST=0.0.0.0
GPIP_LISTEN_PORT=5050
```

**Flags**

```bash
gpip -host 0.0.0.0 -port 5050
```

## Build

Provided build script, `build.sh` will run tests, compile application and finally build a docker image.

### Usage

```bash
Syntax:
# Build binary. Will default to Linux.
./build.sh --version <version>

# Build binary for MacOS.
./build.sh --version <version> --platform darwin

# Build for binary for linux and build docker image.
./build.sh --version <version> --docker
```

## Deployment

`gpip` can be hosted in a matter of ways. It has been tested on:

* Kubernetes (AKS)
* Azure WebApp for Containers

Any platform that can pass on source ip, either through the use of headers `Forwarded`, `X-Forwarded-For`, `X-Real-IP`,
or some other form of manipulating the remote address of the request will do.

### Kubernetes

Manifests for kubernetes deployment (deployment, service and ingress) are provided in the repository. The manifests can be found in `deployments/kubernetes`.

This setup assumes an ingress controller, like `ingress-nginx` is deployed already.

```sh
kubectl create namespace gpip

kubectl apply -f deployments/kubernetes/deployment.yaml -n gpip
kubectl apply -f deployments/kubernetes/service.yaml -n gpip
# Assuming an ingress controller like ingress-nginx is deployed.
kubectl apply -f deployments/kubernetes/ingress.yaml -n gpip
```

As an example, to install the `ingress-nginx` ingress controller do the following:
```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

# The option --set controller.service.externalTrafficPolicy=Local
# ensures that client IP will be forwarded.
helm install nginx-ingress ingress-nginx/ingress-nginx \
    --namespace <namespace> \
    --set controller.replicaCount=2 \
    --set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set controller.service.externalTrafficPolicy=Local
```

If not an ingress controller will be used, but instead the service `.spec.type` is set to `LoadBalancer`, the service `.spec.externalTrafficPolicy` must be set to `Local` to forward the client IP properly.

## Additional credits

Inspiration to intercepting and logging request was inspired/taken from a reply
by user nemith at this [reddit](https://www.reddit.com/r/golang/comments/7p35s4/how_do_i_get_the_response_status_for_my_middleware/) post,
and a reply by huangapple at this [Stack Overflow](https://stackoverflow.com/questions/53272536/how-do-i-get-response-statuscode-in-golang-middleware) post.
So credit to those posts and posters.
