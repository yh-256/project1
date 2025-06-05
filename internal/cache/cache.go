package cache

// EmbeddingCache defines the interface for caching image embeddings.
type EmbeddingCache interface {
	Get(imageID string) ([]float32, error) // Assuming embeddings are float32 slices
	Set(imageID string, embeddings []float32) error
}
