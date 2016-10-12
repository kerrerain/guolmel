package imap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/mxk/go-imap/imap"
	"os"
	"time"
)

type ImapDialer interface {
	DialTLS() (*imap.Client, error)
}

type ImapDialerBasic struct{}

func (i ImapDialerBasic) DialTLS() (*imap.Client, error) {
	if environmentVariablesNotSet() {
		return nil, errors.New("Please set the GUOLMEL_MAIL_{SERVER, USER, PASSWORD}" +
			"environment variables.")
	}

	fmt.Println("Dialing mail server: ", os.Getenv(MAIL_SERVER))

	imapClient, dialError := imap.DialTLS(os.Getenv(MAIL_SERVER),
		&tls.Config{})

	if dialError != nil {
		return nil, dialError
	}

	defer imapClient.Logout(30 * time.Second)

	fmt.Println("Server says hello:", imapClient.Data[0].Info)
	imapClient.Data = nil

	if imapClient.Caps["STARTTLS"] {
		imapClient.StartTLS(nil)
	}

	if imapClient.State() == imap.Login {
		imapClient.Login(os.Getenv(MAIL_USER), os.Getenv(MAIL_PASSWORD))
	}

	imapClient.Select("INBOX", false)

	return imapClient, nil
}

func environmentVariablesNotSet() bool {
	return os.Getenv("GUOLMEL_MAIL_SERVER") == "" ||
		os.Getenv("GUOLMEL_MAIL_USER") == "" ||
		os.Getenv("GUOLMEL_MAIL_PASSWORD") == ""
}
