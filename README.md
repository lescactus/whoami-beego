Who am I ?
==================

A Go web app that display sample information from the visitor (IP, location, browser HTTP headers ...): **[whoami](https://whoami-beego.alexasr.tk/)**.
It thr first attempt of rewriting of one of my previous mini project **[whoami-go](https://github.com/lescactus/whoami-go)** with [beego](https://beego.me/)

This app is strongly inspired by **[ifconfig.me](http://ifconfig.me)**


Use it now
----------
::

```sh
# Install requirements
$ go get "github.com/astaxie/beego" "github.com/ghodss/yaml"

# Run it:
$ bee run
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.10.0
2020/05/03 18:50:29 INFO     ▶ 0001 Using 'whoami-beego' as 'appname'
2020/05/03 18:50:29 INFO     ▶ 0002 Initializing watcher...
github.com/lescactus/whoami-beego/routers
2020/05/03 18:50:31 SUCCESS  ▶ 0003 Built Successfully!
2020/05/03 18:50:31 INFO     ▶ 0004 Restarting 'whoami-beego'...
2020/05/03 18:50:31 SUCCESS  ▶ 0005 './whoami-beego' is running...
2020/05/03 18:50:31.385 [I] [asm_amd64.s:1373]  http server Running on http://:8080
```

Now point your browser at http://127.0.0.1:8080

### Docker
**whoami** can easily be dockerized and is shipped with a ``Dockerfile``.

By default, the container will expose port 8080, so change this within the ``Dockerfile`` if necessary. When ready, simply use the ``Dockerfile`` to build the image.

```sh
cd app
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
