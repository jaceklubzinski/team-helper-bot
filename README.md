# team-helper-bot
![Go](https://github.com/jaceklubzinski/team-helper-bot/workflows/Go/badge.svg?branch=master)
![golangci-lint](https://github.com/jaceklubzinski/team-helper-bot/workflows/golangci-lint/badge.svg?branch=master)
![security scan](https://github.com/jaceklubzinski/team-helper-bot/workflows/security%20scan/badge.svg?branch=master)
![docker build](https://github.com/jaceklubzinski/team-helper-bot/workflows/docker%20build/badge.svg?branch=latest)

# Deployment
Public docker image available on docker hub https://hub.docker.com/repository/docker/jlubzinski/team-helper-bot

docker compose [deployments/docker-compose.yml](deployments/docker-compose.yml) requires environment variables to work
```
# required
HELPERBOT_SLACK_AUTH_TOKEN:
```
# Supported slack command
```
```
