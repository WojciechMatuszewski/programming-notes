# Ultimate docker

## Running Our First Containers

- the `docker stop` command sends a signal to the container. That signal can be ignored. After 10 seconds
  the container is killed (the `kill` command).

- the `docker kill` just kills the container.

## Background containers

- `docker attach` allows you to "be" inside the container

- remember that `CTRL + C` will kill the container if you are inside the container.

- `docker -it` flag will allow you to **interact with `/bin/bash` of that container**.

## Images, Layers & Container Images

- `layers` are usually parts of your application, like the server, your application dependencies.

- `layers` are read only

- `images` are like **classes**, they have **files and metadata**

- `layers` are like inheritance

- `containers` are like **instances** of classes. They are also **read only**.

- since **`containers` are read only**, to create a new container you have to **create a new `image`**.

## Building images interactively

- you do not have to create `Dockerfile` if you do not want to

- you can use `docker diff` and `docker commit` to build an image, you probably want to add a name to the image though.

- `docker diff` is quite nice, it lists the changes from the original container.

## First `Dockerfile`

- `docker` uses cache system so that continuos builds are much faster

- the cache is based on commands. If `docker` sees that instruction did not change, it will be skipped

- you can add `--no-cache` flag to skip the cache (think picking up image after X amount of time)

- the `RUN` command uses the shell that is present on the image. You can also specify the exact command using
  `["", "..']` notation.

## `CMD` and `Entrypoint`

- `CMD` allows you to **run command** when your container spins up

- the `Entrypoint` allows you to **prepend the comand** that you specify with `docker run` with the contents of `Entrypoint`.

  1. `ENTRYPOINT ["figlet"]`
  2. When run `docker run IMAGE_NAME something`
  3. It's the same as `docker run IMAGE_NAME figlet something`

- you can overwrite the `RUN` with just specifying the command: `docker run IMAGE_NAME my_command_that_overwritten`

- to overwrite the `Entrypoint` with the special `--entrypoint` flag

- you can use `CMD` and `Entrypoint` together. The `CMD` acts as a _default value_ that gets prepended with `Entrypoint`

## Copying files

- `COPY` command looks at the files we are copying. This allows it to **skip the cache if the files changed**.
  This makes it so that `--no-cache` flag is not necessary for `COPY` command.

- `ADD` command **does a bit more than `COPY`**, but the instructor recommends sticking with `COPY` command.

## Building images exercise

- another reason for using the `json` syntax for `RUN` command is that the `bin/sh` that runs your program will not pass the signals to that program. You can end up with a container where `ctrl-c` does nothing.
  Instead, you can use the `json` syntax and run the program directly, without `bin/sh`

  ```docker
  RUN ./go_program
  # vs
  RUN ["./go_program"]
  ```

- you can use `exec` bash command to skip the `json` notation.

- **watch out when you use `COPY` with directories**. By default **the contents of the directory is getting copied**.

  ```docker
  COPY src .
  ```

  Now the `/` of the docker image will contain the **contents** of `src` directory.
  To fix this use:

  ```docker
  COPY src src
  ```

## Reducing image size

- you cannot reduce the image size by deleting unwanted dependencies. This does not make sense as **every line within `Dockerfile` creates a layer**. That _layer_ cannot have negative value, the files that you removed will be hidden from the final result, but the image size will not change.

What you **should avoid**:

- _layer collapsing_. This is the notation with `&&` or `\` and multiple commands

  ```docker
  RUN apt-get install .. \
    && apt-get install ..
  ```

  This is pretty bad since chaning 1 character in this whole sequence will bust the cache for that line.

- _outside builds_. This is where you build (or compile) your program on your machine. This is basically loosing everything _Docker_ is about.

What you **should do**. You should prefer **multi-stage builds**.

## Tips and tricks

- you can reduce amount of layers by composing commands using `apt-get dep1 dep2 dep3` syntax. **Only do this after you are sure that you are done with working on the image**. Otherwise you will be busting the cache, a lot.

- **copy `package.json` and install modules first**. This way you can be sure that whenever you only change application code, the modules layer is already cached.

  ```docker
  COPY . . # cache busted pretty much always
  RUN npm install
  ```

  ***

  ```docker
  COPY package.json /
  RUN npm install
  COPY . . # cache busted here, but that's ok because the steps below are pretty fast
  EXPOSE PORT
  CMD ...
  ```

- use `ENV` to get rid of repetition while specifing version names.

  ```docker
  ENV NODE_ENV = 10
  RUN apt-get http://..$NODE_ENV \
      ...
  ```

## Building images with what we learned so far

- remember that _PATH_ within `COPY --from=ALIAS PATH TO` needs to be the path of the workdir. In case of _golang_ image, it's the _/go_ directory.

- some of complied languages like _Golang_ need external dependencies (`dynamic linking`). You can use `golang:alpine` image to build your program and then use the `alpine` itself to run it. Much better than trying to fiddle with `scratch`

  ```docker
  FROM golang:alpine AS builder
  RUN go build main.go

  FROM alpine
  COPY --from=builder /go/main .
  CMD ["./main"]
  ```

## Naming & Inspecting Containers

- container can **only have one unique name**

- by default, `docker` is specifying random names for containers for us.

- use `--name` flag to name the containers. Use `docker start` / `docker stop` for controling the lifecycle.

- you can use `docker inspect` to see the configuration of the container. Use `jq` to parse the `inspect` json.
  Eg. if you want `Created` field

  ```bash
  docker inspect NAME | jq .[0].Created
  ```

  (you might need to use "" if you are using `zsh`)

  You could also use `--format` flag with `docker inspect`. The `--format` flag uses `go templates`.

  ```bash
  docker inspect IMAGE --format "{{ .Name }}"
  ```

## Labels

- a **container can have multiple labels**

- to label images use `-l` flag.

  ```bash
  docker run -l tag=value IMAGE
  ```

## Getting inside a container

- normally you would not need to this, but in reality you might have to

- use the `exec` command to execute command. Combine it with `-ti` flags

  ```bash
  docker exec -ti NAME WHAT (eg.bash)
  ```

- you do not have to use `-ti` if you do not need interactive mode. You can also just run the command itself

  ```bash
  docker exec NAME WHAT (eg.bash or ps)
  ```

- you can export the whole file system of the container

  ```bash
  docker export NAME > where
  ```

- if a container is crashing instantly, overwrite it's `entrypoint` and use the `-ti` flag.

  ```
  docker run -ti --entrypoint sh NAME
  ```

  This way you can get the shell inside the container, and hopefully debug the issue

## Limiting resources

- only useful is you are not using orchestrator like k8s.

- use can limit the memory using `--memory` flag

- you can also limit CPU.

## Networking Basics

- use `EXPOSE` keyword to well, expose ports from your application

- this is basically port mapping. Since we are out of IPs, we have to map ports.

- you can quickly lookup the port mappings you can use the `docker port`

  ```bash
  docker port $(docker ps -lq) OPTIONAL_PORT
  ```

  The `lq` stands for `last` and `quick`.

- the _port mapping_ can be done when you run the container.

  ```bash
  docker run -p HOST_MACHINE_PORT:DOCKER_IMAGE_EXPOSED_PORT DOCKER_IMAGE
  ```

## Container Network Drivers

1. The `bridge` driver
   Think **2 network cards connected togher**. One from the host, one on the container.

2. The `none` driver
   Container will be **disconnected from the network**. You would use it whenever you want your container to not have network access.

3. The `host` driver
   The container will **use the network stack of the host**. This is for maximum performance, normally you would not use it.

4. The `container` driver
   The container will **reuse the network stack of particular container**. This allows for really fast comunication between containers.

## Container Network Model

- think of _wifi_ with a name that you can connect to.

- there is a DNS system built-in. You can assign the DNS name by using `--name` parameter

- you can put 2 containers together in the same network (eg. db and your application) and then use the friendly `name`s.

  ```bash
  docker network create NAME
  docker run -d --net NAME --name FIRST_CONTAINER_NAME IMAGE
  docker run -d --net NAME --name SECOND_CONTAINER_NAME SECOND_IMAGE
  ```

  Now the first and the second container can communicate with each other.

Remember when you tried to run the `amazon/dynamodb-local` image and your golang backend togher? You can name the database whatever you want,
it does not have to be `dynamodb`. Just make sure that they are in the same network and have names.

## Service Discovery With Containers

- use `-P` to just expose all ports. This way you do not have to use mappings with `-p HOST_PORT:CONTAINER_PORT`.
  While the `-P` does not allow you to control the mapping, it's pretty good since you do not have to know which port is exposed by the container.

- use `--net-alias` to assign **only DNS name not the DNS name AND container name** to the container. This is **useful when you want to run
  containers with the same DNS name in different networks**

  ```bash
  docker network create dev
  docker network create prod

  docker run -d --net dev --name redis redis // OK
  docker run -d --net prod --name redis redis // FAIL, container with a name `redis` already exists

  docker run -d --net prod --net-alias redis redis // OK
  ```

- the `--name` or `--net-alias` **does not work with DEFAULT BRIDGE NETWORK**. You **have to create your own network**.

## Local Development with Compose

- sometimes you want to make changes on the host machine (code changes) and you want your docker container to automatically pick it up (probably by some kind of framework that you are using which supports hot-reloading). For this **use volumes**.

- to use volumes you can specify the `-v` flag. The syntax is `HOST:CONTAINER_DIR`

  ```bash
  docker run -d -v $PWD:/src -P NAME
  ```

  The `$PWD` would be the location of source code you copied using `COPY` command

- the files **are not copied**. We are using internal OS mechanics

## More on volumes

- one gotcha is that **if the volume you are mounting is not empty** it will **overwrite the destination**. You can literally overwrite `/usr` or other important folders.

- you can expose `volume`s from the `Dockerfile` directly through `VOLUME` command.

## Compose For Development Stack

- each command inside the given `service` correspond to the `Docker` cli-commands

- by default `docker-compose` creates network for you

- `docker-compose` has built in support for scaling:

  ```bash
  docker-compose up -d --scale SERVICE_NAME=NUM
  ```

## Recap

- use `JSON` notation within the `CMD` to run stuff in the environment when you do not have `bash/sh` installed

- use `&&` to collapse layers. It's very important to ensure that every command is valid.

- use can use `\` to split longer commands (presumably chained with `&&`)

- you **cannot** `COPY` files outside our _build context_

- `WORKDIR` changes the directory for all subsequent comands. Remember that **the `RUN` command** does **not have any "memory"**.
  As in what you did in the previous command has no effect on the next one.
