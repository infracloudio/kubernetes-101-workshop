# Building a container for the app

## App

- A simple go app which just says hello

## Dockerfile


## Building and running container locally

docker build -t app .

docker run -it --rm -p 8080:8080 app