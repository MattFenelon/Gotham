package domainservices

var storer EventStorer

func Configure(es EventStorer) {
	storer = es
}

func ProcessCommand(command *CreateComicCommand) {
	AddComic(command, storer)
}
