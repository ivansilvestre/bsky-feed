package algorithms

import (
	"bsky/pkg/auth"
	"bsky/pkg/feed"
	"github.com/pemistahl/lingua-go"
	"gorm.io/gorm"
)

const SpanishUri = "spanish"

func Spanish(auth *auth.AuthConfig, db *gorm.DB, params feed.QueryParams) (string, []feed.SkeletonItem) {
	language := lingua.Spanish.String()
	params.Language = &language

	cursor, posts := feed.GetFeed(db, params)

	var result = make([]feed.SkeletonItem, 0)
	for _, post := range posts {
		result = append(result, feed.SkeletonItem{Post: post})
	}

	return cursor, result
}
