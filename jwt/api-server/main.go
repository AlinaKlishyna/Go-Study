package main

import (
	"log"

	"github.com/AlinaKlishyna/Go-Study.git/jwt/core"
	"github.com/AlinaKlishyna/Go-Study.git/jwt/core/database"
	"github.com/AlinaKlishyna/Go-Study.git/jwt/internal/api"
)

func main() {
	database.SetupDB()
	domain, err := core.NewDomain()
	if err != nil {
		log.Fatal(err)
	}

	router := api.NewRouter(domain)
	router.Run(":8080")
}
