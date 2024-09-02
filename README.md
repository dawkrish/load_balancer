# load_balancer

From [here](https://codingchallenges.fyi/challenges/challenge-load-balancer)

This challenge is to build your own application layer load balancer.

A load balancer sits in front of a group of servers and routes client requests across all of the servers that are capable of fulfilling those requests. The intention is to minimise response time and maximise utilisation whilst ensuring that no server is overloaded. If a server goes offline the load balance redirects the traffic to the remaining servers and when a new server is added it automatically starts sending requests to it.

Load balancers can work at different levels of the OSI seven-layer network model for example most cloud providers offer application load balancers (layer seven) and network load balancers (layer four). We’re going to focus on a layer seven - application load balancer, which will route HTTP requests from clients to a pool of HTTP servers.
