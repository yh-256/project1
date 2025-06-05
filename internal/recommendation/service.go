package recommendation

import (
	"fmt"
	// "pixiv-tag-reco-service/internal/tagging" // May be needed later for actual tag recall
)

// ArtworkDB is an interface for fetching artwork data.
// This would typically interact with a database like CloudSQL.
type ArtworkDB interface {
	GetArtworksByTags(tags []string, limit int) ([]Artwork, error)
	GetUserFavoriteTags(userID string) ([]string, error) // Example method
}

// BigQueryService is an interface for collaborative filtering via BigQuery.
type BigQueryService interface {
	GetCollaborativeFilterRecommendations(userID string, limit int) ([]Artwork, error)
}

// Service provides artwork recommendation functionality.
type Service struct {
	artworkDB ArtworkDB       // For tag-based recall
	bqService BigQueryService // For collaborative filtering
}

// NewService creates a new Recommendation Service.
// For now, dependencies can be nil as we are stubbing.
func NewService(adb ArtworkDB, bqs BigQueryService) *Service {
	return &Service{artworkDB: adb, bqService: bqs}
}

// GetRecommendationsForUser generates recommendations for a given user.
// This is a stub implementation focusing on tag recall first.
func (s *Service) GetRecommendationsForUser(req Request) (*Response, error) {
	var recommendedArtworks []Artwork

	// --- Stub for Tag Recall ---
	// 1. Get user's favorite tags (placeholder)
	// In a real system, this might come from user profile, interaction history etc.
	userFavoriteTags := []string{"fantasy", "adventure"} // Mock tags
	if s.artworkDB != nil {
		// favTags, err := s.artworkDB.GetUserFavoriteTags(req.UserID)
		// if err == nil && len(favTags) > 0 {
		//  userFavoriteTags = favTags
		// } else if err != nil {
		//  fmt.Printf("Error fetching user favorite tags: %v (using defaults)\n", err)
		// }
		// For stub, let's assume we always try to use the DB if present
		fmt.Printf("Attempting to use artworkDB for user %s (stub)\n", req.UserID)
	}

	fmt.Printf("Using favorite tags for user %s: %v (stub)\n", req.UserID, userFavoriteTags)

	// 2. Recall artworks based on these tags (placeholder)
	if s.artworkDB != nil {
		// For now, even if artworkDB is present, we return mock data to simplify.
		// The actual call would be:
		// artworksFromTags, err := s.artworkDB.GetArtworksByTags(userFavoriteTags, 10) // Get 10 items
		// if err != nil {
		//  return nil, fmt.Errorf("failed to get artworks by tags: %w", err)
		// }
		// recommendedArtworks = append(recommendedArtworks, artworksFromTags...)
		// For stub:
		fmt.Println("artworkDB present, but returning mock artworks for tag recall (stub)")
	}

	// Add some mock artworks for tag recall part
	//nolint:gomnd
	mockTagRecallArtworks := []Artwork{
		{ID: "art001", Title: "Journey to the Mystical Mountain", ImageURL: "/img/art001.jpg", Tags: []string{"fantasy", "landscape"}, Score: 0.8, Source: "tag_recall_stub"},
		{ID: "art002", Title: "Warrior of Light", ImageURL: "/img/art002.jpg", Tags: []string{"fantasy", "character"}, Score: 0.75, Source: "tag_recall_stub"},
	}
	recommendedArtworks = append(recommendedArtworks, mockTagRecallArtworks...)

	// --- End of Stub for Tag Recall ---

	// (BigQuery collaborative filtering will be added in the next step)

	if len(recommendedArtworks) == 0 {
		// Fallback if no recommendations found (e.g., for new users or rare tags)
		//nolint:gomnd
		recommendedArtworks = append(recommendedArtworks, Artwork{
			ID: "fallback001", Title: "Popular This Week", ImageURL: "/img/fallback001.jpg", Tags: []string{"popular"}, Score: 0.5, Source: "fallback_stub",
		})
	}

	return &Response{Artworks: recommendedArtworks}, nil
}
