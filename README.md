[![Go Report Card](https://goreportcard.com/badge/github.com/hitanshu-mehta/reaction-timer)](https://goreportcard.com/report/github.com/hitanshu-mehta/reaction-timer)

## Motivation
The purpose to make this game is to learn primarily about Microservices architecture, gRPC and Go. In addition to that I also learned about Docker, Kubernetes and Prometheus.

## Introduction
This game will show circle and square of different sizes and user have to click them as soon as possible. If user is clicking quickly (i.e having a good reaction time) then size of the shapes will decrease and vice-versa.

## Architecture
![](assets/architecture.png)

### BFF (backend for frontend): 

Receives request from the ui and route to respective microservice.

### highscore:

Set and get highscore.

### gameenigne:

Set the current score of the user and returns size of the shape according to the performation.
If user is performing well then size will be smaller and vice versa.

## Future Plan:

Integrate Thanos to have a global view.