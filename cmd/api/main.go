package main

import (
	"encoding/json" // Was missing in the prompt's main.go, but needed by recommendationHandler's Fprintf
	"log"
	"net/http"
	"pixiv-tag-reco-service/internal/recommendation" // New import
	"pixiv-tag-reco-service/internal/tagging"
)

var tagService *tagging.Service
var recoService *recommendation.Service // New global service variable

// recommendationHandler handles requests for artwork recommendations.
func recommendationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id query parameter", http.StatusBadRequest)
		return
	}

	// This was the placeholder from previous step, now we use the service.
	// fmt.Fprintf(w, "Recommendations endpoint hit for user_id: %s! Placeholder response.", userID)

	req := recommendation.Request{UserID: userID} // Renamed RecommendationRequest to Request
	resp, err := recoService.GetRecommendationsForUser(req)
	if err != nil {
		log.Printf("Error getting recommendations for user %s: %v", userID, err)
		http.Error(w, "Failed to get recommendations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode recommendation response", http.StatusInternalServerError)
		return
	}
}

func tagsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	dummyReq := tagging.ImageTaggingRequest{ImageID: "dummy_image_123"}

	resp, err := tagService.GetTopNTagsForImage(dummyReq)
	if err != nil {
		log.Printf("Error getting tags: %v. Returning mock tags as fallback.", err)
		//nolint:gomnd
		mockTags := []tagging.Tag{
			{Name: "Illustration", Confidence: 0.95},
			{Name: "Anime Style", Confidence: 0.88},
		}
		resp = &tagging.ImageTaggingResponse{Tags: mockTags}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	tagService = tagging.NewService(nil, nil)         // Corrected: Use NewService
	recoService = recommendation.NewService(nil, nil) // Initialize new service

	http.HandleFunc("/v1/tags", tagsHandler)
	http.HandleFunc("/v1/recommendations", recommendationHandler)

	log.Println("Starting server on :8080...")
	//nolint:gosec // G114: Use of net/http serve function that has no support for setting ReadHeaderTimeout
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s", err) // Corrected: No newline
	}
}
