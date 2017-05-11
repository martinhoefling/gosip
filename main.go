package main

import (
	"fmt"
	"github.com/kolo/xmlrpc"
	"os"
	"strings"
)

type SMSQuery struct {
	RemoteUri string
	TOS       string
	Content   string
}

func main() {
	user := os.Getenv("SIPGATE_USERNAME")
	pass := os.Getenv("SIPGATE_PASSWORD")
	url := fmt.Sprintf("https://%s:%s@api.sipgate.net/RPC2", user, pass)

	text := strings.Join(os.Args[2:], " ")
	client, _ := xmlrpc.NewClient(url, nil)
	remoteUri := fmt.Sprintf("sip:%s@sipgate.net", os.Args[1])

	query := SMSQuery{remoteUri, "text", text}

	result := struct {
		SessionID    string
		StatusCode   int
		StatusString string
	}{}

	fmt.Printf("Will send message '%s' to '%s'\n", text, remoteUri)

	err := client.Call("samurai.SessionInitiate", query, &result)
	if err != nil {
		panic(err)
	}
	if result.StatusCode != 200 {
		fmt.Printf("Session: %s\n", result.SessionID)
		fmt.Printf("StatusCode: %d\n", result.StatusCode)
		fmt.Printf("StatusString: %s\n", result.StatusString)
		os.Exit(1)
	}
	fmt.Print("Message successfully sent")
}
