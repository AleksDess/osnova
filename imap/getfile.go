package imap

import (
	"io"
	"io/ioutil"
	"log"
	"mime"
	"strings"

	goimap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	fmail "github.com/emersion/go-message/mail"
)

func GetImapFiles() (ok bool, data []byte, fnm string, err error) {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(ImapClient, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(ImapFiles, ImapFilesPass); err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	if mbox.Messages == 0 {
		ok = false
		return
	}

	// Get the last 4 messages
	seqset := new(goimap.SeqSet)
	seqset.AddRange(mbox.Messages, mbox.Messages)

	messages := make(chan *goimap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []goimap.FetchItem{"BODY[]"}, messages)
	}()

	for msg := range messages {

		r := msg.GetBody(&goimap.BodySectionName{})
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		// Create a new mail reader
		mr, err := fmail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
		}

		// Process each message's part
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			switch p.Header.(type) {
			case *fmail.AttachmentHeader:

				disp, params, err := mime.ParseMediaType(p.Header.Get("Content-Disposition"))
				if err != nil {
					log.Fatal(err)
				}
				if disp == "attachment" {
					filename := params["filename"]

					r := strings.Split(filename, " ")

					for _, i := range r {
						dec := new(mime.WordDecoder)
						decoded, err := dec.Decode(i)
						if err != nil {
							log.Fatal(err)
						}
						fnm += decoded
					}

				}
				data, _ = ioutil.ReadAll(p.Body)
			}
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")

	////////////////////////////////////////////
	// удаление
	item := goimap.FormatFlagsOp(goimap.AddFlags, true)
	flags := []interface{}{goimap.DeletedFlag}
	err = c.Store(seqset, item, flags, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Подтверждение удаления письма
	if err := c.Expunge(nil); err != nil {
		log.Fatal(err)
	}
	/////////////////////////////////////////////
	ok = true
	return
}
