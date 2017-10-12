build: 
	go build -o vigilant

run: build
	sudo ./vigilant

arm-build:
	GOOS=linux GOARCH=arm GOARM=5 go build

build-and-ship: arm-build
	scp camera pi@raspberrypi.local:~/camera
