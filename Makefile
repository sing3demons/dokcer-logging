down:
	docker-compose down
	docker compose -f docker-compose-grafana.yml down
	docker-compose -f docker-compose-fluent-bit.yml down

start:
	docker compose -f docker-compose-grafana.yml up -d
	docker-compose -f docker-compose-fluent-bit.yml up -d
	docker-compose up -d