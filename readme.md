# Onetimesecret

This is a simple service that stores and finds your secret.
Small but powerfull service - does not have any unnesseccery dependencies.

# Hightground overview

The service has 2 endpoints

|Method|Endpoint|Parameters|Response|Description|
|------|--------------|------------------------------------|---------------------------------|------------------------------------------|
|POST  |/create-secret|`{Message: string,Password: string}`|`{Entry:UUID,Expires:Datetime}}`|Saves a secret in DB for 20min             |
|POST  |/secrets      |`{Hash:string,SecretId:UUID}`       |`{Message:string}`              |Extracts and deletes the secret from the DB|

# How to run locally

The easiest way to run this locally is to have docker installed. The server assumes there is a
redis instance running and reachable by `redisdb:6379`. If you prefer to run a local redis server
you can configure your `/etc/hosts` ( or `C:\Windows\System32\drivers\etc\hosts` ) with
```
127.0.0.1 redisdb
```


## Run with docker

* Run docker compose `docker-compose up`

## Run with VSCode

* Open the project as a dev container
* Run `go run .` in the terminal

# Precaution

Although possible to run this service in production it is strongly discouraged. Redis is infamous for being easy pick for hackers.