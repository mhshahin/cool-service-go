# Cool Service - Go

The "Cool Service" enables authenticated and validated users to access a user list, while authenticated admin users can additionally add new users.

User authentication is implemented using JWT (JSON Web Tokens), and access validation is enforced through the use of OPA (Open Policy Agent).
The service utilizes a PostgreSQL database as its data store and employs Helm for deploying both the service and the Open Policy Agent (OPA) service.

## Installation

To run this service, please follow these instructions:

1. Modify the settings within the `Makefile` to align with your specific environment.

2. Deploy a PostgreSQL database either on your cluster or use an existing one. Optionally, you can choose to deploy PostgreSQL elsewhere. You can use the provided `Makefile`` to deploy a PostgreSQL database on your cluster.

3. To execute the necessary database migrations for the service, you must install the `migrate` command on your local machine. You can do this with the following command:

   ```bash
   go get -u -d github.com/golang-migrate/migrate/cmd/migrate
   ```

4. Create a database and run the migrations by running the `make migrate-up` command from the root directory of the project. To create the database, use the following SQL command:

   ```sql
   CREATE DATABASE swisscom;
   ```

   Then, execute the `make migrate-up` command. This will create a `users` table within the `swisscom` database.

5. Update the values in the value files located in the deployment directory to suit your environment.

6. Utilize the `deploy.sh` script within the `deployment/opa` directory to deploy the OPA service.

7. Finally, employ the `deploy.sh` script located in the `deployment/cool_service` directory to deploy and run the service on your cluster.