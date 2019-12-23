package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/dealers"
	"github.com/gorilla/mux"
)

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func Ok(w http.ResponseWriter, body interface{}) {
	_ = json.NewEncoder(w).Encode(body)
}

const addr = ":8080"
const pgHost = "localhost"
const pgDatabase = "dimon_sales_web_core_api_development"

func main() {
	psqlInfo := fmt.Sprintf("host=%s dbname=%s sslmode=disable", pgHost, pgDatabase)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = db.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	dealersStorage := dealers.NewStorage(db)

	router := mux.NewRouter()

	router.HandleFunc("/dealers", func(w http.ResponseWriter, r *http.Request) {
		dealersList, err := dealersStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"dealers": dealersList,
		})
	}).Methods(http.MethodGet)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Listening on address:", addr)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
