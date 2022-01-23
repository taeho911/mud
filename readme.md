# MUD
Project is a playground to develop anything Taeho Kim wants.

## Tech Stack
### Database
  * MongoDB

### Backend
  * Go

### Frontend
  * React JS

### Extra
  * Docker
  * Docker-compose
  * jenkins
  * AWS EC2
  * Amazon Linux 2

## Directory Structure
```
root:
  back:
    agent:              Database agent module
    errors:             Property module to store error codes
    handler:            HTTP handler function module
    model:              Database model module
    router:             HTTP router module
    utils:              Extra util module
  batch:                Batch scripts for inserting dummy data
  ec2:                  Memo for configuring ec2 instance
  env:
    - env.docker.sh:    Environment variables needed for deployment on docker
    - env.local.sh:     Environment variables needed for deployment on local
  front:
    src:
      conponents:       React components
      context:          React contexts
      media:            Media files
      styles:           CSS files
  - docker-compose.yml
  - Jenkinsfile
  - mud_local.sh:       Shell script to build / deploy on local
```

## Architecture

## Functionality
  * Sign up
  * Sign in
  * Session management
  * Deleting user account
  * Money management

## How to Build
There are several environment variables for running application. You can set those by executing **_source env/env.*.sh_** before build.

If you want to change specific variables temporarily, export those environment variables in advance like below. It will preoccupy those environment variables.
```
$ export DB_CONTAINER_NAME=my_database
$ export DB_PORT=27018
$ source ./env/env.docker.sh
```

This behavior is same for both local build and docker build.

### Local
```
$ source env/env.local.sh
$ ./mud_local.sh build

OR

$ source env/env.local.sh
$ cd back && go intall
$ cd ../front && npm run build
```

### Docker
```
$ source env/env.docker.sh
$ docker-compose build
```

## How to Test
```
$ source env/env.docker.sh
$ cd back
$ go test ./...

OR

$ source env/env.docker.sh
$ export BACK_TARGET=test
$ export BACK_IMAGE=${BACK_IMAGE}_test
$ docker-compose build backend
$ docker-compose up -d database
$ docker-compose run backend
```

## How to Deploy
### Local
```
$ source env/env.local.sh
$ ./mud_local.sh start

OR

$ source env/env.local.sh
$ mud
$ cd && npm start
```

### Docker
```
$ source env/env.docker.sh
$ docker-compose up
```
