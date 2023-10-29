NAMESPACE = mhshahin

DB_USERNAME = postgres
DB_PASSWORD = roIoPj6Hja
DB_HOST = localhost
DB_PORT = 5432
DB_NAME = swisscom

serve:
	go run ./main.go serve

config:
	cp config.yaml.example config.yaml

# Change the database connection string according to your own credentials
migrate-up:
	migrate -path ./database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path ./database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

generate-secret:
	openssl rand -hex 32

helm-debug-cool-service:
	helm template coolservice deployment/cool_service -f deployment/cool_service/values.yaml --debug

helm-debug-opa:
	helm template opa deployment/opa -f deployment/opa/values.yaml --debug
	
helm-install-postgres:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	kubectl apply -f deployment/postgres/postgres-pv.yaml
	kubectl apply -f deployment/postgres/postgres-pvc.yaml
	helm install postgresql bitnami/postgresql --set primary.persistence.existingClaim=postgresql-pv-claim --set volumePermissions.enabled=true --namespace $(NAMESPACE)

helm-install:
	helm install opa deployment/opa -f deployment/opa/values.yaml --namespace $(NAMESPACE)
	helm install coolservice deployment/cool_service -f deployment/cool_service/values.yaml --namespace $(NAMESPACE)