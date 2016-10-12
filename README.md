# guolmel
Simple budget management based on a webmail.

## Godep

go get godep
$GOPATH/bin/godep save


## Imap configuration
Export these environment variables:
* GUOLMEL_IMAP_SERVER
* GUOLMEL_IMAP_USER
* GUOLMEL_IMAP_PASSWORD

For example, for a gmail account:
* export GUOLMEL_IMAP_SERVER="imap.gmail.com"
* export GUOLMEL_IMAP_USER="myuser@gmail.com"
* export GUOLMEL_IMAP_PASSWORD="******"

## Smtp configuration