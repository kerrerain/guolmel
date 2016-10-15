package imap

import (
	"errors"
	"github.com/magleff/guolmel/models"
	"github.com/mxk/go-imap/imap"
	"time"
)

func CurrentBudget() (*models.Budget, error) {
	var (
		imapCmd       *imap.Command
		currentBudget *models.Budget
	)

	c, _ := Dial()
	defer c.Logout(30 * time.Second)

	if c.Caps["STARTTLS"] {
		c.StartTLS(nil)
	}

	Login(c)

	c.Select("INBOX", false)

	imapCmd, _ = imap.Wait(c.Search("SUBJECT", c.Quote(models.BUDGET_STATE_FLAG)))

	results := imapCmd.Data[0].SearchResults()

	if len(results) == 0 {
		return nil, errors.New("No active budget for the moment.")
	}

	set, _ := imap.NewSeqSet("")
	set.AddNum(results[len(results)-1])

	imapCmd, _ = c.Fetch(set, "RFC822.HEADER", "BODY[1]")

	for imapCmd.InProgress() {
		// Wait for the next response (no timeout)
		c.Recv(-1)

		if len(imapCmd.Data) > 0 {
			body := imap.AsBytes(imapCmd.Data[0].MessageInfo().Attrs["BODY[1]"])
			currentBudget = models.BudgetFromString(string(body))
		}

		imapCmd.Data = nil
	}

	return currentBudget, nil
}
