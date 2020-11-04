
package lib
import (
	Photo "../models/photo"
	Slack "../models/slack"
	DeleteService "../services"
	PostService "../services"
)

func Delete() {

	photo := Photo.Photo{
		1,
		1,
		"accusamus beatae ad facilis cum similique qui sunt",
		"https://via.placeholder.com/600/92c952",
		"https://via.placeholder.com/150/92c952",
	}
	
	slackMsg := Slack.SlackMsg{
		"New photo deleted title : " + photo.Title,
	}

	DeleteService.DeletePhoto(photo)

	PostService.PostToSlack(slackMsg.SlackMsgToJson())
}
