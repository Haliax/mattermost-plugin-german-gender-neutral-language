package main

import (
	"regexp"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// MessageWillBePosted is invoked when a message is posted by a user before it is committed to the
// database. If you also want to act on edited posts, see MessageWillBeUpdated. Return values
// should be the modified post or nil if rejected and an explanation for the user.
//
// If you don't need to modify or reject posts, use MessageHasBeenPosted instead.
//
// Note that this method will be called for posts created by plugins, including the plugin that created the post.
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	configuration := p.getConfiguration()
	if configuration.disabled {
		return post, ""
	}

	// Always allow posts by the demo plugin user and demo plugin bot.
	if post.UserId == p.botID || post.UserId == configuration.demoUserID {
		return post, ""
	}

	// Process the post message to escape * within words containing "*innen" or "*in"
	post.Message = escapeAsteriskInWord(post.Message)
	return post, ""
}

// MessageWillBeUpdated is invoked when a message is updated by a user before it is committed to
// the database. If you also want to act on new posts, see MessageWillBePosted. Return values
// should be the modified post or nil if rejected and an explanation for the user. On rejection,
// the post will be kept in its previous state.
//
// If you don't need to modify or rejected updated posts, use MessageHasBeenUpdated instead.
//
// Note that this method will be called for posts updated by plugins, including the plugin that
// updated the post.
func (p *Plugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	configuration := p.getConfiguration()
	if configuration.disabled {
		return newPost, ""
	}

	// Process the new post message to escape * within words containing "*innen" or "*in"
	newPost.Message = escapeAsteriskInWord(newPost.Message)
	return newPost, ""
}

func escapeAsteriskInWord(message string) string {
	// Regular expression to find unescaped '*' in words containing '*innen' or '*in'
	re := regexp.MustCompile(`\b(\w+)\*(in(?:nen)?)\b`)

	// Replace the unescaped '*' with '\*'
	return re.ReplaceAllString(message, `$1\*$2`)
}
