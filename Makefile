build:
	# build services image
	cd services/highscore && echo "In services/highscore" \
		docker build -t highscore:v0 .

	cd services/gameengine && echo "In services/gameengine" \
		docker build -t gameengine:v0 .

	docker build -f services/bff/Dockerfile . -t bff:v0

run-highscore:
	docker run highscore:v0

run-gameengine:
	docker run gameengine:v0

run-bff:
	docker run bff:v0