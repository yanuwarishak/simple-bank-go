### How to generate code

- Generate SQL Crud with sqlc:

    ```bash
    make sqlc
    ```

- Create DB mock with gomock:

    ```bash
    make mock
    ```

- Create a new DB migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### How to run

- Run server:

    ```bash
    make server
    ```