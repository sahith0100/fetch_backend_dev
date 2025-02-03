
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor/api/handlers"
    "receipt-processor/internal/storage"
)

func main() {

    store := storage.NewMemoryStorage()
 
    h := handlers.NewHandlers(store)
    
    r := mux.NewRouter()
    
    r.HandleFunc("/receipts/process", h.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", h.GetPoints).Methods("GET")
    
    r.Use(loggingMiddleware)
    
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}