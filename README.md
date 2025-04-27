# foodieeoo

foodieeoo is a Food Ordering Web server application written in GoLang.


## Prerequisites

*   [Go](https://golang.org/dl/) installed (version 1.16 or later recommended).

## Setup

1.  **Clone the repository (if you haven't already):**
    ```bash
    git clone <your-repository-url>
    cd <repository-directory>
    ```

2.  **Install Dependencies:**
    To let Go manage dependencies based on imports, run:
    ```bash
    go mod tidy
    ```
    This command ensures `go.mod` file matches the packages used in the code and downloads them.

3.  **(Optional) Vendor Dependencies:**
    To include dependencies directly in your project repository for offline builds, run:
    ```bash
    go mod vendor
    ```
    This will create a `vendor` directory containing all necessary packages.

4.  **Configure:**
    Create a `config.yml` from `sample_config.yml` file and set the desired configuration values.

## Running the Script

To run the main script, execute:

```bash
go run main.go
```

If you vendored dependencies (Step 3), you might need to build or run using the `-mod=vendor` flag, although `go run` often detects the vendor directory automatically:

```bash
go run -mod=vendor main.go
# or build first
go build -mod=vendor
./<executable-name> # e.g., ./transaction-tracker
```


## Running the DB Migrations

1. Install goose to handle migrations for the service (https://github.com/pressly/goose)

2. To create a new Migration file, run the following command:

    ```goose -dir db/migrations create create_new_table sql```

3. To check the status of migrations, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" status```

4. To upgrade the DB, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" up```

5. To downgrade the DB, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" down```

