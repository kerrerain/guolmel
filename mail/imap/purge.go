package imap

import (
	"github.com/magleff/guolmel/models"
	"github.com/mxk/go-imap/imap"
	"time"
)

func PurgeBudgetStates() error {
	var imapCmd *imap.Command

	c, _ := Dial()
	defer c.Logout(30 * time.Second)

	if c.Caps["STARTTLS"] {
		c.StartTLS(nil)
	}

	Login(c)

	c.Select("INBOX", false)

	imapCmd, _ = imap.Wait(c.Search("SUBJECT", c.Quote(models.BUDGET_STATE_FLAG)))

	set, _ := imap.NewSeqSet("")
	set.AddNum(imapCmd.Data[0].SearchResults()...)

	c.Store(set, "+FLAGS", imap.NewFlagSet(`\Deleted`))

	return nil
}
