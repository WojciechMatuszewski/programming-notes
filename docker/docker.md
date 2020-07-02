# Docker stuff

## Docker DNS

When running multiple containers, maybe using `docker-compose`,
you probably want to connect to some container from other container. This is where **docker DNS** comes in.

See, docker has a DNS built in, it's based on the **container name** (your containers should be in the same network).

I had a problem where I could not connect my `golang` API container with the `dynamodb/local` container.

The `docker-compose` file had these definitions:

```yml
version: "3"
api: ...
dynamo: ...
```

Since, by default, `docker-compose` creates a default network, I though that by using `localhost:8000` I would be able to connect to the dynamodb no problem.

Well, that `localhost:8000` is relative to each container. We should connect to other containers by their name

```yml
version: "3"
api:
  container_name: api
dynamo:
  container_name: dynamo
```

Then I could easily connect to `http://dynamo:8000` since I'm using the default, built-in DNS :)
