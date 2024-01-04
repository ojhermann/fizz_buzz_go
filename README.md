# fizz_buzz_go

This repo contains code in [Go](https://go.dev/) representing [fizz buzz](https://en.wikipedia.org/wiki/Fizz_buzz).

I like [Tom Scott's take on fizz buzz](https://www.youtube.com/watch?v=QPZ0pIK_wsc) and thought it would be fun to
create a direct implementation of the game and then make marginal changes to modify the code in ways that make it better
represent production-level code I would like to see in a Go application.

I'll track this iterative process via [releases](https://github.com/ojhermann/fizz_buzz_go/releases), including a
corresponding `README`.

**Table of Contents**

- [Releases](#releases)
- [Docker stuff](#docker-stuff)

## Releases

- [v0.0.0](README/v0-0-0.md)
- [v0.1.0](README/v0-1-0.md)
- [v0.2.0](README/v0-2-0.md)
- [v0.3.0](README/v0-3-0.md)
- [v0.4.0](README/v0-4-0.md)

## Docker stuff

- `./docker_fizz_buzz_go_development.sh`: run this if you want to develop with feedback from a Docker container e.g. run
  tests in the container to evaluate changes to source code (maybe you don't want to set up Go locally)
    - it will spin up a service named `fizz_buzz_go_development`
    - this service can be accessed from the terminal via `docker compose exec fizz_buzz_go_development bash`; some IDEs
      also offer support for working with containers
      e.g. [Goland](https://blog.jetbrains.com/go/2020/05/06/debugging-a-go-application-inside-a-docker-container/)
    - shutting down the service can be done with `ctrl+c`
    - once the service is shut
      down, [docker compose down](https://docs.docker.com/engine/reference/commandline/compose_down/) is called
- `./docker_fizz_buzz_go_test.sh`: this creates an image and container, runs tests, and then
  calls [docker compose down](https://docs.docker.com/engine/reference/commandline/compose_down/)
    - really only useful for deployment tests; otherwise, you'll save a lot of time and hassle
      using `fbg_development_container`
