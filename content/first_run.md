# First run

Run the nginx container image

```
$ kubectl run nginx --image nginx

```

And then make it available outside the cluster

```
$ kubectl expose deployment nginx --port=80 --target-port=8000
```