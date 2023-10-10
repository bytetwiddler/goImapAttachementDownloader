package main

/*********************************************************
 * using Brian Leishman's go-imap to download            *
 * attachments from an imap server and place them        *
 * in cwd/files/{from address}/                          *
 * Thank you Brian                                       *
 *                                                       *
 * Brian's pgk: http://github.com/BrianLeishman/go-imap/ *
 *********************************************************/

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/BrianLeishman/go-imap"
)

func check(err error) {
	if err != nil {
		log.Println(err.Error)
	}
}

func mkDir(dirName string) {
	if _, statErr := os.Stat(dirName); statErr != nil {
		mkDirErr := os.MkdirAll(dirName, os.ModePerm)
		if mkDirErr != nil {
			panic(mkDirErr)
		}
	}
}

var (
	user   string
	pass   string
	server string
	port   int
)

func init() {
	u, userb := os.LookupEnv("IMAP_USER")
	p, passb := os.LookupEnv("IMAP_PASS")
	s, serverb := os.LookupEnv("IMAP_SERVER")
	pt, serverp := os.LookupEnv("IMAP_PORT")
	if !userb || !passb || !serverb || !serverp {
		fmt.Fprintf(os.Stderr, "Please set required environment variables IMAP_USER , IMAP_PASS, IMAP_SERVER, IMAP_PORT\n")
		os.Exit(1)
	}
	user = u
	pass = p
	server = s
	p64, _ := strconv.ParseInt(pt, 10, 64)
	port = int(p64)
}

func main() {

	// Defaults to false. This package level option turns on or off debugging output, essentially.
	// If verbose is set to true, then every command, and every response, is printed,
	// along with other things like error messages (before the retry limit is reached)
	imap.Verbose = false

	// Defaults to 10. Certain functions retry; like the login function, and the new connection function.
	// If a retried function fails, the connection will be closed, then the program sleeps for an increasing amount of time,
	// creates a new connection instance internally, selects the same folder, and retries the failed command(s).
	// You can check out github.com/StirlingMarketingGroup/go-retry for the retry implementation being used
	imap.RetryCount = 3

	// Create a new instance of the IMAP connection you want to use
	im, err := imap.New(user, pass, server, port)
	check(err)
	defer im.Close()

	// Folders now contains a string slice of all the folder names on the connection
	folders, err := im.GetFolders()
	check(err)

	// Now we can loop through those folders
	for _, f := range folders {

		err = im.SelectFolder(f)
		check(err)

		uids, err := im.GetUIDs("ALL")
		check(err)

		emails, err := im.GetEmails(uids...)
		check(err)

		if len(emails) != 0 {
			for _, e := range emails {
				log.Printf("\n----\nFrom: %v\n", e.From.String())
				log.Printf("Subject: %v\n", e.Subject)
				a := e.Attachments
				for _, f := range a {
					log.Printf("name: %v\ntype: %v\n", f.Name, f.MimeType)
					if len(f.Name) > 0 {
						dirName := fmt.Sprintf("files/%v", e.From.String())
						mkDir(dirName)
						path := fmt.Sprintf("%v/%v", dirName, f.Name)
						log.Printf("path: %v\n", path)
						err := os.WriteFile(path, f.Content, 0644)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}
}
