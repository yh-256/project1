package cache

import "fmt"

// MockEmbeddingCache is a mock implementation of EmbeddingCache for testing.
type MockEmbeddingCache struct {
	GetFunc func(imageID string) ([]float32, error)
	SetFunc func(imageID string, embeddings []float32) error
}

func (m *MockEmbeddingCache) Get(imageID string) ([]float32, error) {
	if m.GetFunc != nil {
		return m.GetFunc(imageID)
	}
	return nil, fmt.Errorf("GetFunc not implemented")
}
func (m *MockEmbeddingCache) Set(imageID string, embeddings []float32) error {
	if m.SetFunc != nil {
		return m.SetFunc(imageID, embeddings)
	}
	return fmt.Errorf("SetFunc not implemented")
}
