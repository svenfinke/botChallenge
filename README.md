# botChallenge

This project is aiming to provide a service that can host various
botChallenges. These botChallenges are different games or tasks that have various
complexities. The service will provide a simple interface to implement a new
game and a REST API that handles the communication to a virtually unlimited count
of players.

# Warning! WIP!
This is still work in progress and in a very early stage of development. All
of the code is subject to change.

# Simulating a bot
If the server is running on port 8000, you can simulate bots like this:

```
Client 1
curl localhost:8000 -H "apikey:hallo123"; for i in seq 0 40; do curl -X POST -H "apikey:hallo123" localhost:8000/api/1; sleep 0.5s; done;

Client 2
curl localhost:8000 -H "apikey:test123"; for i in seq 0 40; do curl -H "apikey:test123" -X POST localhost:8000/api/2; sleep 0.5s; done;
```

# Using the Docker container
The repo contains 2 different services in the compose that can be used:
1. The server
The server can be started with `docker-compose up server`. This will boot a
server on port 8000 for localhost. Now you can connect bots as usual.
2. Build the binary
You can also build the binary with `docker-compose up go_build`. This will build
the botChallenge binary in the src folder.
**Beware:** On the MAC the binary will be build for linux by default, as the VM
for Docker is running Linux. See [this link](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04) to find the correct env vars
for your setup.
