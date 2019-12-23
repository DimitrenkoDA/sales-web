package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/salemaps"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/statuses"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/subs"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/salemans"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/products"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/product_groups"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/pricelist"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/deals"

	_ "github.com/lib/pq"

	"github.com/dimitrenkoda/sales-web/core-api/pkg/dealers"
	"github.com/gorilla/mux"
)

func UnprocessableEntity(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func Ok(w http.ResponseWriter, body interface{}) {
	if body == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	dealsStorage := deals.NewStorage(db)
	pricelistStorage := pricelist.NewStorage(db)
	productGroupStorage := product_groups.NewStorage(db)
	productsStorage := products.NewStorage(db)
	salemansStorage := salemans.NewStorage(db)
	subsStorage := subs.NewStorage(db)
	statusesStorage := statuses.NewStorage(db)
	salemapsStorage := salemaps.NewStorage(db)

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

	router.HandleFunc("/dealers", func(w http.ResponseWriter, r *http.Request) {
		var dealer dealers.Dealer

		err := json.NewDecoder(r.Body).Decode(&dealer)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = dealersStorage.Create(r.Context(), dealer)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, dealer)
	}).Methods(http.MethodPost)
	router.HandleFunc("/dealers/{dealer_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealerID, err := strconv.ParseUint(args["dealer_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var dealer dealers.Dealer

		err = dealersStorage.Find(r.Context(), dealerID, &dealer)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, dealer)
	}).Methods(http.MethodGet)
	router.HandleFunc("/dealers/{dealer_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealerID, err := strconv.ParseUint(args["dealer_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var dealer dealers.Dealer

		err = json.NewDecoder(r.Body).Decode(&dealer)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = dealersStorage.Update(r.Context(), dealerID, dealer)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, dealer)
	}).Methods(http.MethodPut)
	router.HandleFunc("/dealers/{dealer_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealerID, err := strconv.ParseUint(args["dealer_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = dealersStorage.Delete(r.Context(), dealerID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

	router.HandleFunc("/deals", func(w http.ResponseWriter, r *http.Request) {
		dealsList, err := dealsStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"deals": dealsList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/pricelist", func(w http.ResponseWriter, r *http.Request) {
		pricelistList, err := pricelistStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"pricelist": pricelistList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/product_groups", func(w http.ResponseWriter, r *http.Request) {
		productGroupList, err := productGroupStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"product_groups": productGroupList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		productsList, err := productsStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"products": productsList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/salemans", func(w http.ResponseWriter, r *http.Request) {
		salemansList, err := salemansStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"salemans": salemansList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/subs", func(w http.ResponseWriter, r *http.Request) {
		subsList, err := subsStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"subs": subsList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/statuses", func(w http.ResponseWriter, r *http.Request) {
		statusesList, err := statusesStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"statuses": statusesList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/salemaps", func(w http.ResponseWriter, r *http.Request) {
		salemapsList, err := salemapsStorage.List(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"salemaps": salemapsList,
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
