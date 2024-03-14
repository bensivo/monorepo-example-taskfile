# Monorepo Example - Taskfile

Packages:
- helloworld-core: 
  - A Golang lib package
  - Contains importable functions, like `Greet(name)`
- helloworld-api: 
  - Golang HTTP API, 
  - Imports helloworld-core
  - Exposes `GET /greetings/{name}` 
- helloworld-app: 
  - Typescript Vite + React web app, calls helloworld-api
  - Served in an NGINX docker container
  - Calls the helloworld-api endpoint on user form submission
- helloworld: 
  - docker-compose package
  - Runs helloworld-api and helloworld-app in a single command


Commands:

Each package contains a file `Taskfile.dev`, which defines commands. You can run any command from the root directory with `task $PACKAGE_NAME:$COMMAND_NAME`. Some examples:
- Build helloworld-api: `task helloworld-api:build`
- Run helloworld-api: `task helloworld-api:run`
- Run helloworld (api and app): `task helloworld:run`