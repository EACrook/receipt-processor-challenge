package store

import (
	"sync"

	"receipt-processor-challenge/models"
)

type Store struct {
	// ensures proper synchronization for concurrent processing -- ensures only one goroutine can access the resource at a time
	mu sync.Mutex
	receipts map[string]models.StandardReceipt
	points map[string]models.Points
}

func NewStore() *Store {
	return &Store{
		receipts: make(map[string]models.StandardReceipt),
		points: make(map[string]models.Points),
	}
}

func (s *Store) AddReceipt(receipt models.StandardReceipt) {
	s.mu.Lock() // prevents other goroutines from accessing the resource
	defer s.mu.Unlock() // allows other goroutines to access the resource
	s.receipts[receipt.ID] = receipt
}

func (s *Store) AddPoints(points models.Points) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.points[points.ID] = points
}

func (s *Store) GetPoints(id string) (models.Points, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	points, exists := s.points[id]
	return points, exists
}