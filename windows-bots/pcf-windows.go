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
		message.Literal(":windows:"),
		message.Literal("Support: <https://goo.gl/oLxC5y>"),
		message.Prefix(
			"bosh: ",
			message.Conditional(
				pairist.WorkingHours("10:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					pairist.PeopleInRole{
						Team: "boshwindows",
						Role: "Interrupt",
						People: map[string]string{
							"Jason Smith":      "U055LUN2J",
							"Garima Sharma":    "U2LAQS9RQ",
							"Simon Jones":      "U0H6HC71B",
							"Ashwin Venkatesh": "U4DF8VBV0",
							"Kenneth DuMez":    "UBKK5GS8M",
							"Gab Satchi":       "U07K08GGP",
							"Malini Valliath":  "UC53SHZ32",
							"Matthew Horan":    "U0563PD3X",
							"Jackson Feeny":    "U7U1BRTR8",
						},
					},
					message.Literal("PM: <@U0MJ7N77U>"),
				),
			),
		),
		message.Prefix(
			"garden: ",
			message.Conditional(
				pairist.WorkingHours("10:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					pairist.PeopleInRole{
						Team: "garden-windows",
						Role: "Interrupt",
						People: map[string]string{
							"Yael":      "UCKDV7ZEK",
							"Amin":      "U078GBTT7",
							"Natalie":   "U2SSN6RQQ",
							"Yechiel":   "UDT3Q8KLH",
							"Sam":       "U15EHJX52",
							"Arjun":     "UBKBNQBBM",
							"Matt":      "U0563PD3X",
							"Genevieve": "U0BSZH3EX",
							"Sophie":    "UF53HEMMW",
						},
					},
					message.Literal("PM: <@U1CNETREU>"),
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
