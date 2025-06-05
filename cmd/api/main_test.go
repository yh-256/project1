package main

import (
	"context" // Added for NewRequestWithContext
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pixiv-tag-reco-service/internal/tagging"
	"reflect"
	"testing"
)

func TestTagsHandler_PostRequest(t *testing.T) {
	tagService = tagging.NewService(nil, nil) // Renamed taggingService to tagService

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/v1/tags", nil) // Used NewRequestWithContext
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tagsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedTags := []tagging.Tag{
		{Name: "Illustration", Confidence: 0.95},
		{Name: "Anime Style", Confidence: 0.88},
		{Name: "Cute", Confidence: 0.75},
		{Name: "Girl", Confidence: 0.92},
		{Name: "Fantasy", Confidence: 0.80},
	}
	expectedResp := tagging.ImageTaggingResponse{Tags: expectedTags}

	var actualResp tagging.ImageTaggingResponse
	if err := json.NewDecoder(rr.Body).Decode(&actualResp); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if !reflect.DeepEqual(actualResp, expectedResp) {
		t.Errorf("handler returned unexpected body: got %+v want %+v",
			actualResp, expectedResp)
	}
}

func TestTagsHandler_WrongMethod(t *testing.T) {
	tagService = tagging.NewService(nil, nil)                                                     // Renamed taggingService to tagService
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/v1/tags", nil) // Used NewRequestWithContext
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tagsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for GET: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
