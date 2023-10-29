serve:
	go run ./main.go serve

config:
	cp config.yaml.example config.yaml

# Change the database connection string according to your own credentials
migrate-up:
	migrate -path ./database/migrations -database "postgres://postgres:roIoPj6Hja@localhost:5433/swisscom?sslmode=disable" up

migrate-down:
	migrate -path ./database/migrations -database "postgres://edward@localhost:5432/swisscom?sslmode=disable" down

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
	helm install postgresql bitnami/postgresql --set primary.persistence.existingClaim=postgresql-pv-claim --set volumePermissions.enabled=true --namespace mhshahin

mocks:
	/Users/edward/go/bin/mockgen -destination=mocks/user_mock.go -package=mocks github.com/mhshahin/cool-service-go/repository/user_repository User



kubectl get secret --namespace mhshahin postgresql -o jsonpath={.data.postgres-password} | base64 -d