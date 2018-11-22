# Pod

# What is a pod

- Set of containers running together - which share port space, network and storage
- Most common pattern is of single container in pod, but Kubernetes managed pod and that in turn manages container

![pod](pod.svg)

# Run the pod

```
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: app
```

# Run a pod with 2 containers

# Init container

# Pod Termination - grace period

# Readiness Probe

# Liveness Probe

# Restart policies

# Delete a pod - what happens

# Resource allocation

# Volumes - Quick overview
