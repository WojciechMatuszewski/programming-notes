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
