package lib

import (
	Photo "../models/photo"
	Slack "../models/slack"
	PostService "../services"
)

func Post() {
	
	photo := Photo.Photo {
		1,
		5001,
		"foto title",
		"https://via.placeholder.com/600/e21cb",
		"https://via.placeholder.com/150/9e66a4",
	}
	
	slackMsg := Slack.SlackMsg{
		"New photo created",
	}

	PostService.PostToServer(photo)
	PostService.PostToSlack(slackMsg.SlackMsgToJson())
	
}
