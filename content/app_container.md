# From app to Kubernetes in 60 seconds

## App

- A simple go app which just says hello

## Dockerfile

- A multistage Dockerfile - first part builds source code into binary

- Second part takes binary from first part and creates an image


## Building and running container locally

```
docker build -t app .

docker run -it --rm -p 8080:8080 app
```

## Pushing image

```
$ docker tag app vishalbiyani/app
$ docker push vishalbiyani/app
```

## Write pod spec

Every Kubernetes object is defined by a standard format 

- The `apiVersion` and `kind` define the kind of object. 
- `Metadata` has all the metadata about that object such as name, labels, annotations etc.
- `Spec` - is the core definition of the object itself and varies from object to object

```
apiVersion: v1
kind: <ObjectKind>
metadata:
  name: <Name>
  labels:
    <key>: <value>
spec:
  name: <Name>
```

So let's define a simple pod spec

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
    image: vishalbiyani/app

```

## Create Service spec

Similarly define a service spec:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: simpleapp
  labels:
    svc: myapp
spec:
  type: LoadBalancer
  ports:
  - port: 80
    targetPort: 8080
  selector:
    app: myapp
```

## Apply specs

```bash
$ kubectl apply -f pod.yaml
$ kubectl apply -f service.yaml
```

## Validate application works

```
$ kubectl get service

```