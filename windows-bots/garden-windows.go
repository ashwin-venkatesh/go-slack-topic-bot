package main

import (
	"log"
	"os"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/dpb587/go-slack-topic-bot/slack"
)

func main() {
	msg, err := message.Join(
		" | ",
		message.Literal(":garden: :windows:"),
		message.Prefix(
			"interrupt: ",
			message.Conditional(
				pairist.WorkingHours("10:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					pairist.PeopleInRole{
						Team: "garden-windows",
						Role: "Interrupt",
						People: map[string]string{
							"Yael":      "UCM54MLUQ",
							"Amin":      "U077YLS56",
							"Natalie":   "U2W83JVH9",
							"Yechiel":   "UDXNS2XHU",
							"Sam":       "U15D0FLTW",
							"Arjun":     "UBP6RMANS",
							"Matt":      "U04G2B2U1",
							"Genevieve": "U1Y770WAU",
							"Sophie":    "UF8DFN9D0",
						},
					},
					message.Literal("PM: <@U1FJWH8HL>"),
				),
			),
		),
	).Message()
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv("SLACK_CHANNEL"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}
