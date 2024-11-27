package main

import (
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
)

// TestEscapeAsteriskInWord tests the escapeAsteriskInWord function
func TestEscapeAsteriskInWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Tester*in", "Tester\\*in"},
		{"Bereits gegenderte Freund\\*in", "Bereits gegenderte Freund\\*in"},
		{"Mehrere Freund*innen und Kollekt*innen", "Mehrere Freund\\*innen und Kollekt\\*innen"},
		{"Nichts zu tun", "Nichts zu tun"},
		{"Freund*innen und Kolleg\\*innen **sollen nicht** interargieren", "Freund\\*innen und Kolleg\\*innen **sollen nicht** interargieren"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := escapeAsteriskInWord(test.input)
			if result != test.expected {
				t.Errorf("unexpected result for %q: got %q, want %q", test.input, result, test.expected)
			}
		})
	}
}

// TestMessageWillBePosted tests the MessageWillBePosted function
func TestMessageWillBePosted(t *testing.T) {
	p := &Plugin{botID: "testBot"}
	post := &model.Post{UserId: "userId", Message: "Freund*in"}

	modifiedPost, _ := p.MessageWillBePosted(nil, post)

	expectedMessage := "Freund\\*in"
	if modifiedPost.Message != expectedMessage {
		t.Errorf("unexpected message: got %q, want %q", modifiedPost.Message, expectedMessage)
	}
}

// TestMessageWillBeUpdated tests the MessageWillBeUpdated function
func TestMessageWillBeUpdated(t *testing.T) {
	p := &Plugin{botID: "testBot"}
	newPost := &model.Post{UserId: "userId", Message: "Freund*innen"}
	oldPost := &model.Post{UserId: "userId", Message: "Old message"}

	modifiedPost, _ := p.MessageWillBeUpdated(nil, newPost, oldPost)

	expectedMessage := "Freund\\*innen"
	if modifiedPost.Message != expectedMessage {
		t.Errorf("unexpected message: got %q, want %q", modifiedPost.Message, expectedMessage)
	}
}
