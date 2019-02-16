# The Binding of Isaac: Afterbirth+ achievement helper [![Build Status](https://travis-ci.org/fuglede/afterbirthachievements.svg?branch=master)](https://travis-ci.org/fuglede/afterbirthachievements)

Simple web application for allowing a given Steam user to follow their progress towards completing achievements in [Afterbirth+](http://store.steampowered.com/app/570660/). The application is currently live at https://tboi.mosstier.com

As anyone familiar with it will undoubtedly notice, this is very strongly inspired by the incredibly useful http://theriebel.de/tboirah/, which aggregates achievements for Rebirth and Afterbirth. 


## Building

To run this application locally, carry out the following steps:
* Rename `config.json.example` to `config.json` and modify it to include the relevant information.
* Then, build it using `go build`.
* After running the application, navigate to `http://localhost:9090` (unless you changed the port in the config).

## Running from Docker Hub

Alternatively, the application can be run directly from Docker Hub through

```
$ docker run -d -p 9090:9090 -e STEAMAPIKEY=9783B7D64EABC6060606060606060666 fuglede/afterbirthachievements
```

where the value of `STEAMAPIKEY` should be replaced with your Steam API key.
