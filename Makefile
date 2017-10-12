dev-env:
	which statik || cd vendor/github.com/rakyll/statik && go install

static:
	statik -src=./public

build: static 
	go build -o vigilantpi

run-%: build
	sudo ./vigilantpi $*

arm-build: static
	GOOS=linux GOARCH=arm GOARM=5 go build -o vigilantpi

build-and-ship: arm-build
	scp vigilantpi pi@raspberrypi.local:~/.
