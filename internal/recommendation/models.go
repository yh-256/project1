package recommendation

// Artwork represents a piece of artwork.
type Artwork struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	Tags     []string `json:"tags,omitempty"`   // Tags associated with the artwork
	Score    float64  `json:"score,omitempty"`  // Recommendation score
	Source   string   `json:"source,omitempty"` // e.g., "tag_recall", "collaborative_filtering"
}

// Request holds parameters for recommendation requests.
// Currently only UserID, but could be extended.
type Request struct {
	UserID string
	// Potentially other parameters like PageToken, PageSize, etc.
}

// Response is the response payload for the recommendation API.
type Response struct {
	Artworks []Artwork `json:"artworks"`
	// Potentially NextPageToken, etc.
}
