# team-helper-bot
![Go](https://github.com/jaceklubzinski/team-helper-bot/workflows/Go/badge.svg?branch=master)
![golangci-lint](https://github.com/jaceklubzinski/team-helper-bot/workflows/golangci-lint/badge.svg?branch=master)
![security scan](https://github.com/jaceklubzinski/team-helper-bot/workflows/security%20scan/badge.svg?branch=master)
![docker build](https://github.com/jaceklubzinski/team-helper-bot/workflows/docker%20build/badge.svg?branch=latest)

Features:
- reply in the thread to defined phrases
- managed list of problems with short solutions
- adds a reaction to greetings
# Deployment
Public docker image available on docker hub https://hub.docker.com/repository/docker/jlubzinski/team-helper-bot

docker compose [deployments/docker-compose.yml](deployments/docker-compose.yml) requires environment variables to work
```
# required
HELPERBOT_SLACK_AUTH_TOKEN:
```
# Supported slack command
`help` help message

`add` add problem with possible solution
```
short: 
@bot add sources.LazyJDBCSource http://github.com

long: 
@bot add "ProxySQL Error: Access denied for user" "recreate container"
```
`list` list all problems and solutions

`del` del single row 
```
@bot del "single row"
```

`fix-all` delete all problems

