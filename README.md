# K8s net test

This project was designed "in the field" at customers sites where network stability and response times were to be tested and analyzed.

## Components

There are two main components a server and a client Pods, the server is deployed with 3 replicas by default and the client with one.

The client will communicate with the server's service and show response times.

Load can be adjusted:
1. The number of subroutines (mimicing clients).
2. The time between each request on each subroutine.

### Server

The server has no real configuration, it will expose port 8080 and reply to /ping with a Json.

### Client

The client is more configurable by changing the environment variables:
- FORK - The number of subroutines the client will use.
- DELAYVAL - The delay between requests in each subroutine.
- DEST - The destination service for which the 3 server Pod are used as endpoints.


## Usage on K8s

deployment is simple:

```
kubectl apply -f yalms/net-test.yml
```

assuming you have an internet connection the images will be downloaded automatically.

The client variables are changeable by modifying the yamls.
