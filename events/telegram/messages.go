package telegram

const (
	msgHelp = `I can save urls(not pages). 
	In order to save the page, just send me a link to it.

	In order to get a random page from your list, send me command /rnd.
	WARNING! Ur page will be removed from list after /rnd command.
	`

	msgHello          = `Greetings, traveller! \n\n` + msgHelp
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "This bitch is empty"
	msgSaved          = "Saved"
	msgAlreadyExists  = "You have already saved this page"
)
