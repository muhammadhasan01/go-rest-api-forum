package main

import (
	"backend-forum/migrations"
)

func main() {
	migrations.Migrate()
	// migrations.CreateAccounts()
}
