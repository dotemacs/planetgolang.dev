package main

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	InitializeDb("file::memory:?cache=shared")

	Db.AutoMigrate(&Post{})
	publishedAt, _ := time.Parse("2006-Jan-02", "2022-Feb-02")
	Db.Create(&Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	})

	var posts []Post
	results := Db.Find(&posts)
	post := posts[0]

	if err := results.Error; err != nil {
		t.Errorf("Error in finding posts: %v", err)
	}

	if results.RowsAffected != 1 || post.Title != "title" {
		t.Errorf("Unexpected post data retrieved: %v", posts)
	}

	if post.FormattedShortPublishedAt() != "02 Feb 22" {
		t.Errorf("Post's short published at date is incorrect: %s", post.FormattedShortPublishedAt())
	}

	if post.FormattedPublishedAt() != "02 Feb 22 00:00 UTC" {
		t.Errorf("Post's formatted published at date is incorrect: %s", post.FormattedPublishedAt())
	}
}
