
package storage

import (
    "sync"
    "errors"
    "github.com/google/uuid"
    "receipt-processor/api/models"
)

type MemoryStorage struct {
    receipts map[string]*models.Receipt
    mu       sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        receipts: make(map[string]*models.Receipt),
    }
}

func (s *MemoryStorage) SaveReceipt(receipt *models.Receipt) string {
    s.mu.Lock()
    defer s.mu.Unlock()

    id := uuid.New().String()
    s.receipts[id] = receipt
    return id
}

func (s *MemoryStorage) GetReceipt(id string) (*models.Receipt, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    receipt, exists := s.receipts[id]
    if !exists {
        return nil, errors.New("receipt not found")
    }
    return receipt, nil
}
