# Containerized WebApp

This is a short Go and Reactjs web app tom find the highest prime number of a given, this structure of this project is inspired 
by [Golang Project Layout](https://github.com/golang-standards/project-layout)

A live version of this app is hosted on [here](http://206.189.31.248:8080/) on digital ocean

## Design Decisions
I decided to go with a concurrency approach, utilising Golang's concurrency support 
I was able to break finding the highest prime number into batches of 10, i.e if an input 50 is supplied,
the last 10 highest number of 50 that is 40-50 will be calculated for prime numbers if no prime number is
found within this range the next 10 highest numbers 30-40 will be calculated and so on.

Using this approach we're taking advantage of Golang's multi-processing ability to solve this problem faster,
when the highest prime number of a particular supplied number is found the highest prime number is indexed in redis database for 
an hour to make future retrievals a lot faster

## Running
The app and its dependencies are managed by docker compose to run the app use the command below it uses docker-compose under the hood
```bash
$ make run
```

## Testing
```bash
$ cd __tests__
$ go test
```

