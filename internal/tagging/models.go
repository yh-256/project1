package tagging

// Tag represents a single tag with its confidence score.
type Tag struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// ImageTaggingRequest represents the request payload for the tagging API.
// It would typically include image data or a URL.
// For now, we'll keep it simple.
type ImageTaggingRequest struct {
	ImageID string `json:"image_id"` // Placeholder for image identifier
}

// ImageTaggingResponse represents the response payload for the tagging API.
type ImageTaggingResponse struct {
	Tags []Tag `json:"tags"`
}
