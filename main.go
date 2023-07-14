package main

import (
	"fmt"
	"osnova/imap"
)

/*
git add .
git commit -m "Your commit message"

git push
*/

func main() {
	fmt.Println(1)
	PrintGetEnv()

	imap.GetImapFiles()
}
