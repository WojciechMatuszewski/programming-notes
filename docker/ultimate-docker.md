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
