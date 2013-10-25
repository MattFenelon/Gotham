package domainservices

import (
	"domain"
)

type EventStorer interface {
	AddEvent(event *domain.ComicAdded)
}
