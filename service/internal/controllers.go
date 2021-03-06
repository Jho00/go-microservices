package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ProduceGetCustomerHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		GetCounterInstance().Inc()
		log.Println(fmt.Sprintf("Customer handler connections amount: %d", GetCounterInstance().Value()))
		channel := make(chan []byte)
		go func() {

			customers, err := json.Marshal(GetCustomers(db))

			if err != nil {
				log.Fatal(err)
				w.Write([]byte(err.Error()))
				return
			}

			channel <- customers
		}()
		customers := <- channel
		close(channel)
		w.Write(customers)
		GetCounterInstance().Dec()
	}
}

func ProduceOrderCustomerHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		GetCounterInstance().Inc()
		log.Println(fmt.Sprintf("Order handler connections amount: %d", GetCounterInstance().Value()))
		channel := make(chan []byte)
		go func() {
			//time.Sleep(20 * time.Second)

			id := r.URL.Query().Get("id")
			if id == "" {
				w.Write([]byte("Please pass id in query param"))
				return
			}

			orders, err := json.Marshal(GetOrders(db, id))

			if err != nil {
				log.Fatal(err)
				w.Write([]byte(err.Error()))
				return
			}

			channel <- orders
		}()
		orders := <- channel
		close(channel)
		w.Write(orders)
		GetCounterInstance().Dec()
	}
}
