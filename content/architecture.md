# Kubernetes Architecture


Source: https://kubernetes.io/docs/concepts/architecture/cloud-controller/
![](./static/pre-ccm-arch.png)

## Master Components

### kube-apiserver

Kubernetes API server is how you interact with Kubernetes. It is stateless component and scales horizontally. When you use the Kubectl CLI - it basically talks to Kubernetes API server

### etcd

Etcd is a key value store  and is where all of data of Kubernetes is stored. For scalability and redundancy you need to deploy it in HA configuration. 

### Kube-scheduler 

This component basically looks at scheduling the pods onto nodes. The criteria used for scheduling could vary but looks at overall resource needs, available resources affinity/anti-affinity considerations etc.

### Kube controller manager

This component is for running various controllers on Kubernetes. Controller is essentially a control loop which is running and making changes so that actual state reflects the desired state. Kubernetes has a few built in controllers:

#### Node controller

#### Replication Controller

#### Endpoints controller 

#### Service accounts and token controller

### Cloud controller manager

This is purely for interfacing with underlying cloud which could be AWS, Google cloud, Azure and so on. This layer abstracts the underlying cloud's specific ways of working and works with rest of Kubernetes.

## Node components

### Kubelet

This agent runs on every node of the cluster. It manages local things such as running containers etc. It does not manage any container which is not scheduled by Kubernetes

### Kube-proxy

Manages the networking part by changing the routing etc. using IP tables.

### Container runtime

One of container runtimes like Docker, Rkt, runc which actually runs container

## Additional Components

### DNS

For routing to services with human readable names, DNS component is needed. KubeDNS is one of earlier components but the newly developed CoreDNS is better. 

There is a DNS specification for Kubernetes: https://github.com/kubernetes/dns/blob/master/docs/specification.md 


### Dashboard

### Resources monitoring

### Cluster logging

