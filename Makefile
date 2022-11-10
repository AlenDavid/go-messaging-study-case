BASE=alen

.PHONY: services
services: consumer producer

.PHONY: consumer
consumer:
	docker build --build-arg SERVICE=consumer -t ${BASE}/hello-consumer .

.PHONY: producer
producer:
	docker build --build-arg SERVICE=producer -t ${BASE}/hello-producer .
