# Deployment

- Deployment creates N number of pods and manages them - such as upgrades, rolling deploys etc.

- Deployment creates a replica set and replica set in turn creates pods.

## Create a deployment

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

- Create deployment

- Show deployments

- Check deployment specs and status fields

- Show events of deployment

- Show how deployment, replicaset and pods are linked with labels


## Update deployment image

- Update using CLI but it can also be done by editing the deployment

- Notice the record - we want to track the changes to deployment

```
$ kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.9.1 --record

```

Check the status of deployment updates:

```
$ kubectl rollout status deployment.v1.apps/nginx-deployment
```

## Rolling back deployment

Let's first check the status of rollout history

```
$ kubectl rollout history deployment.v1.apps/nginx-deployment
```

```
$ kubectl rollout undo deployment.v1.apps/nginx-deployment

```

- Also you can rollback to a specific version by using `--to-revision`

## Scale a deployment

Scaling a deployment:

```
$ kubectl scale deployment.v1.apps/nginx-deployment --replicas=10

```

If you want autoscaling (Which requires few other things to work)

```
$ kubectl autoscale deployment.v1.apps/nginx-deployment --min=10 --max=15 --cpu-percent=80

```

## Fail a deployment Upgrade

On purpose let's give a wrong tag in update and look at events

```
$ kubectl set image deployment.v1.apps/nginx-deployment nginx=nginx:1.912

```

## Rolling updates

### Max Unavailable & Max Surge 

- Change max unavailable and surge and then upgrade again with right image tag