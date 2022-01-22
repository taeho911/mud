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

## How to Build

## How to Deploy
