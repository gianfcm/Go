package models

import (
	"encoding/json"
	"log"
)

type SlackMsg struct {
	Text string `json:"text"`
}

func (slackmsg SlackMsg) SlackMsgToJson() []byte {
	slackToJson, err := json.Marshal(slackmsg)
	if err != nil {
		log.Fatalf("Error on slackMsgToJson: %v", err)
	}
	return slackToJson
}