# homeAssistance


to run it on rasphi4 (/linux/arm/v8 enviroment)

docker run --priviledged --rm tonistiigi/btnfmt --install all

before initialize.






step1. 
	cd docker
	docker build -t homeassist . (for rasphi, docker build --platform linux/amd64)
	docker run -t homeassis homeassist
	docker exec -it homeassis ./startup.sh

step2. 
	cd ../controller
	go build . (MAKEFILE will be added)
	 
