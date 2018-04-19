package messages

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	sq "github.com/nleiva/lab-inventory/sql"
	"github.com/nlopes/slack"
)

var (
	setUserStmt       *sql.Stmt
	setSwStmt         *sql.Stmt
	getUserStm        *sql.Stmt
	err               error
	acceptedGreetings = map[string]bool{
		"what's up?": true,
		"hey!":       true,
		"yo":         true,
	}
	acceptedHowAreYou = map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
	}
	acceptedCommand = map[string]bool{
		"reserve": true,
		"release": true,
		"check":   true,
	}
)

const (
	setUserQ = "UPDATE device_table SET user=? WHERE node = ?"
	setSwQ   = "UPDATE device_table SET sw_image=? WHERE node = ?"
	getUserQ = "SELECT user FROM device_table WHERE node = ?"
)

// Listen receives Slack events and triggers actions.
func Listen(db *sql.DB, token string) chan []string {
	// Prepared MySQL statements. The actual SQL queries are defined as constants.
	if setUserStmt, err = db.Prepare(setUserQ); err != nil {
		log.Fatal(fmt.Errorf("mysql: prepare set user: %v", err))
	}
	if setSwStmt, err = db.Prepare(setSwQ); err != nil {
		log.Fatal(fmt.Errorf("mysql: prepare set software: %v", err))
	}
	if getUserStm, err = db.Prepare(getUserQ); err != nil {
		log.Fatal(fmt.Errorf("mysql: prepare get user: %v", err))
	}

	s := make(chan []string)
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	go func(chan []string) {
	Loop:
		for {
			select {
			case msg := <-rtm.IncomingEvents:
				// fmt.Println("Event Received: ")
				switch ev := msg.Data.(type) {
				case *slack.HelloEvent:
					// Ignore hello
				case *slack.ConnectedEvent:
					// fmt.Println("Infos:", ev.Info)
					// fmt.Println("Connection counter:", ev.ConnectionCount)
					rtm.SendMessage(rtm.NewOutgoingMessage("Hello again!", "#go-testing"))
				case *slack.MessageEvent:
					// fmt.Printf("Message: %v\n", ev)
					info := rtm.GetInfo()
					botuser := info.User.ID
					// Set a prefix that should be met in order to warrant a response from us
					prefix := fmt.Sprintf("<@%s> ", botuser)
					// If the original message wasn’t posted by our bot AND it contains our
					// prefix @botuser, then we’ll respond to the channel.
					// if ev.User != botuser && strings.HasPrefix(ev.Text, prefix) {
					//	rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))
					// }
					if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
						respond(rtm, ev, prefix, s)
					}
				case *slack.PresenceChangeEvent:
					// fmt.Printf("Presence Change: %v\n", ev)
				case *slack.LatencyReport:
					// fmt.Printf("Current latency: %v\n", ev.Value)
				case *slack.RTMError:
					fmt.Printf("Error: %s\n", ev.Error())

				case *slack.InvalidAuthEvent:
					fmt.Println("Invalid credentials")
					close(s)
					break Loop

				default:
					// Ignore other events..
				}
			}
		}
	}(s)
	return s
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string, ch chan []string) {
	var response string
	text := msg.Text
	if len(text) < 1 {
		response = "Do you want to say something to me?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		return
	}
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	array := strings.Split(text, " ")

	if acceptedGreetings[text] {
		response = "What's up buddy!?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedCommand[array[0]] {
		if len(array) < 2 {
			response = "Error: Need $command $argument(s)"
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			return
		}
		execCommand(rtm, msg, array)
		ch <- array
	} else {
		response = "What!?... "
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

func execCommand(rtm *slack.RTM, msg *slack.MessageEvent, array []string) {
	var response string
	command := array[0]
	device := array[1]
	switch command {
	case "reserve":
		response = fmt.Sprintf("Reserving %s", device)
		if len(array) < 3 {
			response = "Error -> reserve $device $user"
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			return
		}
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		err = sq.SetUser(setUserStmt, device, array[2])
		if err != nil {
			response = "Error: " + err.Error()
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			return
		}
	case "release":
		response = fmt.Sprintf("Releasing %s", device)
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		err = sq.SetUser(setUserStmt, device, "none")
		if err != nil {
			response = "Error: " + err.Error()
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			return
		}
	case "check":
		user, err := sq.GetUser(getUserStm, device)
		if err != nil {
			response = "Error: " + err.Error()
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
			return
		}
		response = fmt.Sprintf("In use by %s", user)
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}
