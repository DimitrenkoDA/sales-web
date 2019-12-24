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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	router.HandleFunc("/deals", func(w http.ResponseWriter, r *http.Request) {
		var deal deals.Deal

		err := json.NewDecoder(r.Body).Decode(&deal)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = dealsStorage.Create(r.Context(), deal)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, deal)
	}).Methods(http.MethodPost)
	router.HandleFunc("/deals/{deal_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealID, err := strconv.ParseUint(args["deal_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var deal deals.Deal

		err = dealsStorage.Find(r.Context(), dealID, &deal)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, deal)
	}).Methods(http.MethodGet)
	router.HandleFunc("/deals/{deal_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealID, err := strconv.ParseUint(args["deal_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var deal deals.Deal

		err = json.NewDecoder(r.Body).Decode(&deal)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = dealsStorage.Update(r.Context(), dealID, deal)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, deal)
	}).Methods(http.MethodPut)
	router.HandleFunc("/deals/{deal_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		dealID, err := strconv.ParseUint(args["deal_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = dealsStorage.Delete(r.Context(), dealID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/pricelist", func(w http.ResponseWriter, r *http.Request) {
		var pricelist pricelist.Pricelist

		err := json.NewDecoder(r.Body).Decode(&pricelist)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = pricelistStorage.Create(r.Context(), pricelist)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, pricelist)
	}).Methods(http.MethodPost)
	router.HandleFunc("/pricelist/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		prodID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var pricelist pricelist.Pricelist

		err = pricelistStorage.Find(r.Context(), prodID, &pricelist)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, pricelist)
	}).Methods(http.MethodGet)
	router.HandleFunc("/pricelist/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		prodID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var pricelist pricelist.Pricelist

		err = json.NewDecoder(r.Body).Decode(&pricelist)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = pricelistStorage.Update(r.Context(), prodID, pricelist)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, pricelist)
	}).Methods(http.MethodPut)
	router.HandleFunc("/pricelist/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		prodID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = pricelistStorage.Delete(r.Context(), prodID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/product_groups", func(w http.ResponseWriter, r *http.Request) {
		var productGroup product_groups.ProductGroup

		err := json.NewDecoder(r.Body).Decode(&productGroup)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = productGroupStorage.Create(r.Context(), productGroup)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, productGroup)
	}).Methods(http.MethodPost)
	router.HandleFunc("/product_groups/{pg_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		PgID, err := strconv.ParseUint(args["pg_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var productGroup product_groups.ProductGroup

		err = productGroupStorage.Find(r.Context(), PgID, &productGroup)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, productGroup)
	}).Methods(http.MethodGet)
	router.HandleFunc("/product_groups/{pg_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		PgID, err := strconv.ParseUint(args["pg_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var productGroup product_groups.ProductGroup

		err = json.NewDecoder(r.Body).Decode(&productGroup)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = productGroupStorage.Update(r.Context(), PgID, productGroup)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, productGroup)
	}).Methods(http.MethodPut)
	router.HandleFunc("/product_groups/{pg_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		PgID, err := strconv.ParseUint(args["pg_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = productGroupStorage.Delete(r.Context(), PgID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		var product products.Product

		err := json.NewDecoder(r.Body).Decode(&product)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = productsStorage.Create(r.Context(), product)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, product)
	}).Methods(http.MethodPost)
	router.HandleFunc("/products/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		ProdID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var product products.Product

		err = productsStorage.Find(r.Context(), ProdID, &product)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, product)
	}).Methods(http.MethodGet)
	router.HandleFunc("/products/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		ProdID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var product products.Product

		err = json.NewDecoder(r.Body).Decode(&product)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = productsStorage.Update(r.Context(), ProdID, product)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, product)
	}).Methods(http.MethodPut)
	router.HandleFunc("/products/{prod_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		ProdID, err := strconv.ParseUint(args["prod_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = productsStorage.Delete(r.Context(), ProdID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/salemans/top5", func(w http.ResponseWriter, r *http.Request) {
		top5List, err := salemansStorage.ShowTop5(r.Context())

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"top5": top5List,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/salemans/unsold", func(w http.ResponseWriter, r *http.Request) {
		salemanName := r.URL.Query().Get("saleman_name")
		leftDate := r.URL.Query().Get("left_date")
		rightDate := r.URL.Query().Get("right_date")

		unsoldList, err := salemansStorage.UnsoldProduct(r.Context(), salemanName, leftDate, rightDate)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, map[string]interface{}{
			"unsold": unsoldList,
		})
	}).Methods(http.MethodGet)

	router.HandleFunc("/salemans", func(w http.ResponseWriter, r *http.Request) {
		var saleman salemans.Saleman

		err := json.NewDecoder(r.Body).Decode(&saleman)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = salemansStorage.Create(r.Context(), saleman)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, saleman)
	}).Methods(http.MethodPost)
	router.HandleFunc("/salemans/{man_code}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemanID, err := strconv.ParseUint(args["man_code"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var saleman salemans.Saleman

		err = salemansStorage.Find(r.Context(), SalemanID, &saleman)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, saleman)
	}).Methods(http.MethodGet)
	router.HandleFunc("/salemans/{man_code}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemanID, err := strconv.ParseUint(args["man_code"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var saleman salemans.Saleman

		err = json.NewDecoder(r.Body).Decode(&saleman)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = salemansStorage.Update(r.Context(), SalemanID, saleman)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, saleman)
	}).Methods(http.MethodPut)
	router.HandleFunc("/salemans/{man_code}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemanID, err := strconv.ParseUint(args["man_code"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = salemansStorage.Delete(r.Context(), SalemanID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/subs", func(w http.ResponseWriter, r *http.Request) {
		var sub subs.Sub

		err := json.NewDecoder(r.Body).Decode(&sub)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = subsStorage.Create(r.Context(), sub)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, sub)
	}).Methods(http.MethodPost)
	router.HandleFunc("/subs/{sub_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SubID, err := strconv.ParseUint(args["sub_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var sub subs.Sub

		err = subsStorage.Find(r.Context(), SubID, &sub)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, sub)
	}).Methods(http.MethodGet)
	router.HandleFunc("/subs/{sub_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemanID, err := strconv.ParseUint(args["sub_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var sub subs.Sub

		err = json.NewDecoder(r.Body).Decode(&sub)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = subsStorage.Update(r.Context(), SalemanID, sub)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, sub)
	}).Methods(http.MethodPut)
	router.HandleFunc("/subs/{sub_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SubID, err := strconv.ParseUint(args["sub_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = subsStorage.Delete(r.Context(), SubID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/statuses", func(w http.ResponseWriter, r *http.Request) {
		var status statuses.Status

		err := json.NewDecoder(r.Body).Decode(&status)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = statusesStorage.Create(r.Context(), status)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, status)
	}).Methods(http.MethodPost)
	router.HandleFunc("/statuses/{status_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		StatusID, err := strconv.ParseUint(args["status_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var status statuses.Status

		err = statusesStorage.Find(r.Context(), StatusID, &status)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, status)
	}).Methods(http.MethodGet)
	router.HandleFunc("/statuses/{status_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		StatusID, err := strconv.ParseUint(args["status_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var status statuses.Status

		err = json.NewDecoder(r.Body).Decode(&status)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = statusesStorage.Update(r.Context(), StatusID, status)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, status)
	}).Methods(http.MethodPut)
	router.HandleFunc("/statuses/{status_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		StatusID, err := strconv.ParseUint(args["status_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = statusesStorage.Delete(r.Context(), StatusID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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

	router.HandleFunc("/salemaps", func(w http.ResponseWriter, r *http.Request) {
		var salemap salemaps.Salemap

		err := json.NewDecoder(r.Body).Decode(&salemap)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = salemapsStorage.Create(r.Context(), salemap)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, salemap)
	}).Methods(http.MethodPost)
	router.HandleFunc("/salemaps/{map_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemapID, err := strconv.ParseUint(args["map_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var salemap salemaps.Salemap

		err = salemapsStorage.Find(r.Context(), SalemapID, &salemap)

		if err == sql.ErrNoRows {
			NotFound(w)
			return
		} else if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, salemap)
	}).Methods(http.MethodGet)
	router.HandleFunc("/salemaps/{map_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemapID, err := strconv.ParseUint(args["map_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		var salemap salemaps.Salemap

		err = json.NewDecoder(r.Body).Decode(&salemap)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}

		err = salemapsStorage.Update(r.Context(), SalemapID, salemap)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, salemap)
	}).Methods(http.MethodPut)
	router.HandleFunc("/salemaps/{map_id}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		SalemapID, err := strconv.ParseUint(args["status_id"], 10, 64)

		if err != nil {
			UnprocessableEntity(w, err)
			return
		}
		err = statusesStorage.Delete(r.Context(), SalemapID)

		if err != nil {
			InternalServerError(w, err)
			return
		}

		Ok(w, nil)
	}).Methods(http.MethodDelete)

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
