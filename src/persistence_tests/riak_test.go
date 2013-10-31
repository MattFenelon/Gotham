package riak_tests

import (
	"bytes"
	"domain"
	"encoding/gob"
	"github.com/mrb/riakpbc"
	// "persistence"
	"testing"
)

// TODO: Rename file?

func TestComicAdded(t *testing.T) { // TODO: Better test name
	client, err := Connect()
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer func() {
		// TODO: Catch panic when close fails.
		client.Close()
	}()

	seriesTitle, _ := domain.NewSeriesTitle("Prophet")
	title, _ := domain.NewBookTitle("Prophet 31")
	expectedEvent := domain.NewComicAdded(domain.NewRandomComicId(), seriesTitle, title)

	storeErr := StoreEvent(client, expectedEvent)
	if storeErr != nil {
		t.Fatalf("%v", err)
	}

	actualEvent, getErr := GetEvent(client, expectedEvent.Id.String())
	if getErr != nil {
		t.Fatalf("%v", getErr)
	}

	if expectedEvent.Equal(actualEvent) == false {
		t.Errorf("Expected:\n%v\nbut was:\n%v", expectedEvent, actualEvent)
	}
	// TODO: Delete bucket at end of test.
}

func Connect() (client *riakpbc.Client, err error) {
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

func StoreEvent(client *riakpbc.Client, event *domain.ComicAdded) error {
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

// TODO: What type should comicId be? What's the correct interface into the domain?
func GetEvent(client *riakpbc.Client, comicId string) (event *domain.ComicAdded, err error) {
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
