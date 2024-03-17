# Monorepo Example - Taskfile

Packages:
- helloworld: 
  - docker-compose package
  - Runs helloworld-api and helloworld-app in a single command
- helloworld-api: 
  - Golang HTTP API, 
  - Exposes `POST /greet` 
- helloworld-app: 
  - Typescript Vite + React web app, calls the `POST /greet` endpoint
- helloworld-proto: 
  - Protobuf type definitions (GreetRequest, GreetResponse)
  - Uses buf.build to compile protos 
- helloworld-proto-go: 
  - Generates golang code for helloworld-proto types
- helloworld-proto-ts: 
  - Generates ts code for helloworld-proto types


Commands:

Each package contains a file `Taskfile.dev`, which defines commands. You can run any command from the root directory with `task $PACKAGE_NAME:$COMMAND_NAME`. Some examples:
- Build helloworld-api: `task helloworld-api:build`
- Run helloworld-api: `task helloworld-api:run`
- Run helloworld (api and app): `task helloworld:run`