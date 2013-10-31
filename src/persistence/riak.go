package persistence

import (
	"bytes"
	"domain"
	"encoding/gob"
	"github.com/mrb/riakpbc"
)

// TODO: Keep connections open for the length of a HTTP request rather than opening and closing
// the connection for each operation.

type RiakEventStore struct {
	cluster  []string
	clientId string
}

func NewRiakEventStore(cluster []string, clientId string) *RiakEventStore {
	return &RiakEventStore{
		cluster:  cluster,
		clientId: clientId}
}

func (r *RiakEventStore) connect() (client *riakpbc.Client, err error) {
	// TODO: Error if Riak database is not running
	// TODO: Network addresses should be configurable
	client = riakpbc.NewClient(r.cluster)
	if err = client.Dial(); err != nil {
		return client, err
	}

	// TODO: Better ClientId
	if _, err = client.SetClientId(r.clientId); err != nil {
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
	if _, err := client.StoreObject("comics", event.Id.String(), buffer.Bytes()); err != nil {
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

	rsp, err := client.FetchObject("comics", comicId)
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

func (r *RiakEventStore) DeleteEvent(comicId string) error {
	client, err := r.connect()
	if err != nil {
		return err
	}
	defer func() {
		// TODO: Catch panic when close fails.
		client.Close()
	}()

	if _, err = client.DeleteObject("comics", comicId); err != nil {
		return err
	}

	return nil
}
