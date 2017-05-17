# go_rage_online
Go application for CPU load testing

This application will host a web server on port 8001 by default. It is designed to be used with a container and AutoScaling.

It is also able to show personalised data in the form of a environment variable. "SPECIAL_DATA". Using this you can set messages in your containers to view special instances.

Example:
```
Starting CPU stress on: 1e7dd605c92d
Specail data is set, value is: TheRandyOne
```
If you want to build this binary and put it in a scratch container you need to build it with some special flags.

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_rage_online .
```
This will ensure that you can run the application inside an empty container by pulling in all dynamic linking.
