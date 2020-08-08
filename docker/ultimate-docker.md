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
