# Docker centralized logging using Fluent Bit, Grafana and Loki

Grafana team has released Loki, which is inspired by Prometheus to solve this issue. So now, we don't need to manage multiple stacks to monitor the running systems like Grafana and Prometheus to monitor and EFK to check the logs.

### Setup

- docker-compose-grafana.yml
- docker-compose-fluent-bit.yml
- fluent-bit.conf
- docker-compose-app.yml

Create network
```
$docker network create loki
```

Create grafana + loki
```
$docker compose -f docker-compose-grafana.yml up -d
```

URL = http://localhost:3000
* user=admin
* password=admin


Create fluent-bit
```
$docker-compose -f docker-compose-fluent-bit.yml up -d
```

Send log from app to fluent-bit
```
$docker-compose -f docker-compose-app.yml up -d
```

List of urls 
* http://localhost:4000
* http://localhost:4000/test


## Config grafana
* Datasource
  * Loki = http://loki:3100
  * Explore with Loki ...

