
package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor/api/models"
    "receipt-processor/internal/storage"
    "receipt-processor/internal/calculator"
)

type Handlers struct {
    storage *storage.MemoryStorage
}

func NewHandlers(storage *storage.MemoryStorage) *Handlers {
    return &Handlers{storage: storage}
}

func (h *Handlers) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt

    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid request body. Please verify input.", http.StatusBadRequest)
        return
    }

    if err := receipt.Validate(); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id := h.storage.SaveReceipt(&receipt)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.ReceiptResponse{ID: id})
}

func (h *Handlers) GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    receipt, err := h.storage.GetReceipt(id)
    if err != nil {
        http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
        return
    }

    points, err := calculator.CalculatePoints(receipt)
    if err != nil {
        http.Error(w, "Error calculating points. Please verify input.", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.PointsResponse{Points: points})
}