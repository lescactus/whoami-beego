Who am I ?
==================

A Go web app that display sample information from the visitor (IP, location, browser HTTP headers ...): **[whoami](https://whoami-beego.alexasr.fr/)**.
It the first attempt of rewriting of one of my previous mini project **[whoami-go](https://github.com/lescactus/whoami-go)** with [beego](https://beego.vip/)

This app is strongly inspired by **[ifconfig.me](http://ifconfig.me)**


Use it now
----------

```sh
# Install via go get and run it
$ go get "github.com/lescactus/whoami-beego"
$ ~/go/bin/whoami-beego
...

# Git clone and build:
$ git clone https://github.com/lescactus/whoami-beego.git
$ go run main.go
...
```

Now point your browser at http://127.0.0.1:8080

### Docker
**whoami** can easily be dockerized and is shipped with a ``Dockerfile``.

By default, the container will expose port 8080, so change this within the ``Dockerfile`` if necessary. When ready, simply use the ``Dockerfile`` to build the image.

```sh
docker build -t whoami .
```
This will create the Docker image.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 80 of the host to port 5000 of the container:

```sh
docker run -d -p 80:8080 --restart="always" --name whoami whoami 
```

Now point your browser at http://127.0.0.1/

Screenshots
-----------
![IP info location](https://i.imgur.com/y1EMwDe.png "IP info location")
***
![Map](https://i.imgur.com/QN4JMiX.png "Map")
***
![Sitemap](https://i.imgur.com/PCyz1qo.png "Site map")
