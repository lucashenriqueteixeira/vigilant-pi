# Vigilant Pi

This project is an attempt to write a software that can search for [cheap] survillence IP cameras on local network,and periodically record their videos and upload them to Google Drive.

It's aimed to run on a Raspberry Pi 3 and with luck, on a Raspberry Pi Zero W.

## Dependencies
- Go >= 1.8.3
- [dep](https://github.com/golang/dep)
- [Nmap](https://nmap.org/)

## Getting started
```sh
# install dependencies first

go get github.com/petersondmg/vigilant-pi

cd $GOPATH/src/github.com/petersondmg/vigilant-pi

# install statik only at first time
make dev-env

# admin interface
make run-admin & 

# watcher
make run-watch
```