package persistence

import (
	"bytes"
	"domain"
	"encoding/gob"
	"github.com/mrb/riakpbc"
)

type RiakEventStore struct {
}

func NewRiakEventStore() *RiakEventStore {
	return &RiakEventStore{}
}

func (r *RiakEventStore) connect() (client *riakpbc.Client, err error) {
	// TODO: Error if Riak database is not running
	// Network addresses should be configurable
	client = riakpbc.NewClient([]string{"127.0.0.1:8080"})
	if err = client.Dial(); err != nil {
		return client, err
	}

	// TODO: Better ClientId
	if _, err = client.SetClientId("test"); err != nil {
		return client, err
	}

	return client, nil
}

func (r *RiakEventStore) AddComic(event *domain.ComicAdded) error {
	client, err := r.connect()
	if err != nil {
		return err
	}
	defer func() {
		// TODO: Catch panic when close fails.
		client.Close()
	}()

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(event); err != nil {
		return err
	}

	// TODO: Where should the bucket name come from?
	// Key change required.
	// Set the content type
	if _, err := client.StoreObject("testbucket", "test", buffer.Bytes()); err != nil {
		return err
	}

	return nil
}

func (r *RiakEventStore) GetEvent(comicId string) (event *domain.ComicAdded, err error) {
	client, err := r.connect()
	if err != nil {
		return nil, err
	}
	defer func() {
		// TODO: Catch panic when close fails.
		client.Close()
	}()

	rsp, err := client.FetchObject("testbucket", "test")
	if err != nil {
		return nil, err
	}
	content := rsp.GetContent()
	value := bytes.NewBuffer(content[0].GetValue())
	dec := gob.NewDecoder(value)

	event = &domain.ComicAdded{}
	err = dec.Decode(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
