package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/joho/godotenv"
)

func main() {
	// Define the execution context, supplying a migration directory
	// and potentially an `atlas.hcl` configuration file using `atlasexec.WithHCL`.
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("../migrations"),
		),
	)
	if err != nil {
		log.Fatalf("failed to load working directory: %v", err)
	}
	// atlasexec works on a temporary directory, so we need to close it
	defer workdir.Close()

	// Initialize the client.
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	fmt.Println("Loading env vars...")
	errorGoDotEnv := godotenv.Load("../.env")
	if errorGoDotEnv != nil {
		log.Fatalf("failed to load env vars: %v", errorGoDotEnv)
	}
	userDatabase := os.Getenv("USER_DATABASE")
	passwordDatabase := os.Getenv("PASSWORD_DATABASE")
	addressDatabase := os.Getenv("ADDRESS_DATABASE")
	portDatabase := os.Getenv("PORT_DATABASE")
	nameDatabase := os.Getenv("NAME_DATABASE")

	// Run `atlas migrate apply` on a SQLite database under /tmp.
	fmt.Println("Running migrations...")
	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: "mysql://" + userDatabase + ":" + passwordDatabase + "@" + addressDatabase + ":" + portDatabase + "/" + nameDatabase,
	})
	if err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	fmt.Printf("Applied %d migrations\n", len(res.Applied))
}
