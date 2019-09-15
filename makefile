build:
	docker build --tag ${DOCKER_REPOSITORY}/${PROJECT}/${SERVICE} .

buildTest:
	docker build --tag ${DOCKER_REPOSITORY}/${PROJECT}/${SERVICE}:test .