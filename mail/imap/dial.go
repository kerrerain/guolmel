package imap

import (
	"errors"
	"fmt"
	"github.com/mxk/go-imap/imap"
	"os"
)

func Dial() (*imap.Client, error) {
	if serverVariableNotSet() {
		return nil, errors.New("Please set the GUOLMEL_IMAP_SERVER " +
			"environment variable.")
	}

	c, err := imap.DialTLS(os.Getenv(SERVER), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Server says hello:", c.Data[0].Info)
	c.Data = nil

	return c, nil
}

func Login(c *imap.Client) error {
	_, err := c.Login(os.Getenv(USER), os.Getenv(PASSWORD))
	return err
}

func serverVariableNotSet() bool {
	return os.Getenv(SERVER) == ""
}

func credentialsVariableNotSet() bool {
	return os.Getenv(USER) == "" ||
		os.Getenv(PASSWORD) == ""
}
