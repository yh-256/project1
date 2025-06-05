package tagging

import (
	"fmt"
	"pixiv-tag-reco-service/internal/cache"
	"testing"
)

func TestTaggingService_GetTopNTagsForImage_NoCLIPModel(t *testing.T) {
	service := NewService(nil, nil)
	req := ImageTaggingRequest{ImageID: "test_image_no_clip"}
	resp, err := service.GetTopNTagsForImage(req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected a response, got nil")
	}
	if len(resp.Tags) != 5 {
		t.Errorf("Expected 5 mock tags, got %d", len(resp.Tags))
	}
}

func TestTaggingService_GetTopNTagsForImage_WithCLIPModel(t *testing.T) {
	mockClip := &MockCLIPModel{}
	service := NewService(mockClip, nil)
	req := ImageTaggingRequest{ImageID: "test_image_with_clip"}
	_, err := service.GetTopNTagsForImage(req)

	if err == nil {
		t.Errorf("Expected an error for 'CLIP model not fully implemented yet', got nil")
	} else if err.Error() != "CLIP model not fully implemented yet" {
		t.Errorf("Expected error 'CLIP model not fully implemented yet', got '%v'", err)
	}
}

func TestTaggingService_GetTopNTagsForImage_CacheHit(t *testing.T) {
	mockCache := &cache.MockEmbeddingCache{
		GetFunc: func(imageID string) ([]float32, error) {
			if imageID == "cached_image" {
				return []float32{0.1, 0.2, 0.3}, nil
			}
			return nil, fmt.Errorf("cache miss")
		},
	}
	service := NewService(nil, mockCache)
	req := ImageTaggingRequest{ImageID: "cached_image"}
	resp, err := service.GetTopNTagsForImage(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(resp.Tags) != 5 {
		t.Errorf("Expected 5 mock tags after cache hit (stub behavior), got %d", len(resp.Tags))
	}
}

func TestTaggingService_GetTopNTagsForImage_CacheMissAndSet(t *testing.T) {
	cacheSetCalled := false
	var setImgID string
	mockCache := &cache.MockEmbeddingCache{
		GetFunc: func(imageID string) ([]float32, error) {
			return nil, fmt.Errorf("cache miss for new_image_for_cache")
		},
		SetFunc: func(imageID string, embeddings []float32) error {
			if imageID == "new_image_for_cache" {
				setImgID = imageID
				cacheSetCalled = true
			}
			return nil
		},
	}
	service := NewService(nil, mockCache)
	req := ImageTaggingRequest{ImageID: "new_image_for_cache"}
	_, err := service.GetTopNTagsForImage(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !cacheSetCalled {
		t.Errorf("Expected cache.Set to be called on cache miss for 'new_image_for_cache'")
	}
	if setImgID != "new_image_for_cache" {
		t.Errorf("Expected cache.Set to be called with imageID 'new_image_for_cache', got '%s'", setImgID)
	}
}
