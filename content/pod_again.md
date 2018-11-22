
# Pod - Deep Dive

## Run a pod with 2 containers

## Init container

- Great for initialization needed before starting the apps

- Init container does not live beyond the initialization process

- Can have more than one init container

- Following pod will not start until the two endpoints are reachable

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
  - name: init-myservice
    image: busybox
    command: ['sh', '-c', 'until nslookup myservice; do echo waiting for myservice; sleep 2; done;']
  - name: init-mydb
    image: busybox
    command: ['sh', '-c', 'until nslookup mydb; do echo waiting for mydb; sleep 2; done;']
```

And the services which will enable above pod's initialization process:

```yaml
kind: Service
apiVersion: v1
metadata:
  name: myservice
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
---
kind: Service
apiVersion: v1
metadata:
  name: mydb
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9377
```

- Use cases
  - Initialization of a secret/auth data from external service
  - Checking an external system is up before starting etc.

## Pod Termination - grace period

- Pre termination hooks allow to cleanup/persist state before exiting

## Health Checks

- Demo app - probes_app scenarios

### Readiness Probe

### Liveness Probe

## Restart policies

Three restart policies:

- Always
- OnFailure
- Never

## Resource allocation

- Requests

- Limits

## Volumes - Quick overview
