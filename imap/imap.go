package imap

import (
	"fmt"
	"io/ioutil"

	"log"
	"net/mail"
	"osnova/logger"
	"strings"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

var Speed_message = ""
var Hourly_report = ""
var Behavior_report = ""
var ImapClient = ""
var ImapFiles = ""
var ImapFilesPass = ""

type MessageSpeed struct {
	Car    string
	Date   string
	Speed  string
	Adress string
	Link   string
}

// Печать элементаMessageSpeed
func (a *MessageSpeed) Print() {
	fmt.Println()
	fmt.Println("MessageSpeed        :")
	fmt.Println("Car                 :", a.Car)
	fmt.Println("Date                :", a.Date)
	fmt.Println("Speed               :", a.Speed)
	fmt.Println("Adress              :", a.Adress)
	fmt.Println("Link                :", a.Link)
}

type ListMessageSpeed []MessageSpeed

func (a *ListMessageSpeed) Print() {
	for _, i := range *a {
		fmt.Printf("i: %v\n", i)
	}
}

func GetImapBox(mailadress, key, box string) (mbox *imap.MailboxStatus, c *client.Client, err error) {

	c, err = client.DialTLS(ImapClient, nil)
	if err != nil {
		return
	}

	if err = c.Login(mailadress, key); err != nil {
		return
	}

	mbox, err = c.Select(box, false)
	if err != nil {
		return
	}

	return
}

func Boltmessage(mbox *imap.MailboxStatus, c *client.Client, deleteemails bool) (numberofmessages int, res ListMessageSpeed) {

	from := uint32(1)
	to := mbox.Messages
	numberofmessages = int(to)

	if mbox.Messages > 10 {
		from = mbox.Messages - 10
		numberofmessages -= 10
	} else {
		return
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	for i := from; i < to; i++ {

		seqset := new(imap.SeqSet)
		//seqset.AddRange(, uint32(i))
		seqset.AddNum(uint32(i))

		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope}

		messages := make(chan *imap.Message, 1)
		done := make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, items, messages)
		}()

		msg := <-messages
		// log.Println("* " + msg.Envelope.MessageId)
		r := msg.GetBody(section)
		if r == nil {
			logger.ErrorLog.Println("Server didn't returned message body")
		}

		if err := <-done; err != nil {
			logger.ErrorLog.Println(err)
		}

		m, err := mail.ReadMessage(r)
		if err != nil {
			logger.ErrorLog.Println(err)
		}

		body, err := ioutil.ReadAll(m.Body)
		if err != nil {
			logger.ErrorLog.Println(err)
		}
		rs := MessageSpeed{}
		s := string(body)
		ss := strings.Split(s, "<br />")

		for _, i := range ss {

			s := i
			s = strings.ReplaceAll(s, "\n", "")
			s = strings.ReplaceAll(s, "\t", "")
			s = strings.ReplaceAll(s, string([]byte{13}), "")
			s = strings.Trim(s, " ")
			if len(s) < 10 {
				continue
			}
			pars := parsestring(s)

			if len(pars) == 4 {
				if pars[0] == "speed" {
					rs.Car = pars[3]
					rs.Speed = pars[2]
					rs.Date = pars[1]
				}
				if pars[0] == "adress" {
					rs.Adress = pars[1]
					rs.Link = pars[2]
				}
			}

		}
		res = append(res, rs)
	}

	if deleteemails {
		for i := from; i < to; i++ {
			////////////////////////////////////////////
			// удаление
			item := imap.FormatFlagsOp(imap.AddFlags, true)
			flags := []interface{}{imap.DeletedFlag}
			err := c.Store(seqset, item, flags, nil)
			if err != nil {
				log.Fatal(err)
			}
			// Подтверждение удаления письма
			if err := c.Expunge(nil); err != nil {
				log.Fatal(err)
			}
			/////////////////////////////////////////////
		}
	}
	return
}

func (a *MessageSpeed) IsCity() bool {
	r := a.Adress
	r = strings.Trim(r, " ")
	s := strings.Split(r, ", ")
	if len(s) == 0 {
		return false
	}
	b := []byte(s[len(s)-1])
	if len(b) != 5 {
		return false
	}
	for _, i := range b {
		fmt.Println(i)
		if i < 48 || i > 57 {
			return false
		}
	}
	return true
}

func parsestring(s string) (r []string) {
	if strings.Contains(s, "[") && strings.Contains(s, "]") {
		// fmt.Println(s)
		b := []rune(s)
		var n1, n2, n3, n4, fl1, fl2 int
		for n, i := range b {
			if i == 91 && fl1 == 0 {
				n1 = n
				fl1++
			}
			if i == 93 && fl2 == 0 {
				n2 = n
				fl2++
			}
			if i == 91 && fl1 != 0 {
				n3 = n
			}
			if i == 93 && fl2 != 0 {
				n4 = n
			}
		}
		r = append(r, "speed")
		r = append(r, string(b[n1+1:n2]))
		r = append(r, string(b[n3+1:n4]))
		r = append(r, string(b[n2+4:n2+12]))
	}
	if strings.Contains(s, "Aдрес") {
		b := []rune(s)
		var n1 int
		for n, i := range b {
			if i == 104 {
				n1 = n
				break
			}
		}
		r = append(r, "adress")
		r = append(r, string(b[:n1]))
		r = append(r, string(b[n1:]))
		r = append(r, "")
	}
	return
}

func Boltfiles() {
	fmt.Println("Start")

	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(ImapClient, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(ImapFiles, ImapFilesPass); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	log.Println("Total messages", to)
	if mbox.Messages > 10 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 10
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)
	items := []imap.FetchItem{imap.FetchEnvelope}

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	log.Println("Last 10 messages:")
	log.Println("Len messages:", len(messages))
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
