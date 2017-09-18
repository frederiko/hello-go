# hellopilmeo

## Building simplehttp
``` 
$ go get github.com/pilmeo/hellopilmeo
If you are running on Linux
$ GOOS=linux GOARCH=amd64 go build .
``` 

## Creating the Docker image
```
$ cd $GOPATH/src/github.com/pilmeo/simplehttp
$ docker build -t fcosta/simplehttp .
```
## Running the container
```
$ docker run -p 8080:8080 fcosta/simplehttp 
$ docker run -p 8080:8080 -d fcosta/simplehttp (in background)
```

