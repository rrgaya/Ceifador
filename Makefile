deploy:
	docker build -t ceifador:latest . && docker tag ceifador us-west1-docker.pkg.dev/conversion-toolkit/toolkitrepo/ceifador:latest && docker push us-west1-docker.pkg.dev/conversion-toolkit/toolkitrepo/ceifador:latest

build:
	docker build -t ceifador:latest . && docker run --restart=always -d ceifador:latest
