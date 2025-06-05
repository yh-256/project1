package tagging

import "fmt"

// MockCLIPModel is a mock implementation of CLIPModel for testing.
type MockCLIPModel struct {
	GetTagsForImageFunc       func(imageData []byte) ([]Tag, error)
	GetEmbeddingsForImageFunc func(imageData []byte) ([]float32, error)
}

func (m *MockCLIPModel) GetTagsForImage(imageData []byte) ([]Tag, error) {
	if m.GetTagsForImageFunc != nil {
		return m.GetTagsForImageFunc(imageData)
	}
	return nil, fmt.Errorf("GetTagsForImageFunc not implemented")
}
func (m *MockCLIPModel) GetEmbeddingsForImage(imageData []byte) ([]float32, error) {
	if m.GetEmbeddingsForImageFunc != nil {
		return m.GetEmbeddingsForImageFunc(imageData)
	}
	return nil, fmt.Errorf("GetEmbeddingsForImageFunc not implemented")
}
