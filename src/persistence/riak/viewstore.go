package riak

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/MattFenelon/riakpbc"
)

type ViewStore struct {
	bucket   string
	cluster  []string
	clientId string
}

func NewViewStore(cluster []string, clientId string) *ViewStore {
	return &ViewStore{
		bucket:   "views",
		cluster:  cluster,
		clientId: clientId,
	}
}

func (v *ViewStore) connect() (client *riakpbc.Client, err error) {
	// TODO: Error if Riak database is not running
	// TODO: Network addresses should be configurable
	client = riakpbc.NewClient(v.cluster)
	if err = client.Dial(); err != nil {
		return client, err
	}

	// TODO: Better ClientId
	if _, err = client.SetClientId(v.clientId); err != nil {
		return client, err
	}

	return client, nil
}

func (v *ViewStore) Store(key string, in interface{}) error {
	client, err := v.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	var sendBuffer bytes.Buffer
	enc := gob.NewEncoder(&sendBuffer)
	if err := enc.Encode(in); err != nil {
		return err
	}
	if _, err := client.StoreObject(v.bucket, key, sendBuffer.Bytes()); err != nil {
		return err
	}

	return nil
}

func (v *ViewStore) Get(key string, out interface{}) error {
	client, err := v.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	rsp, err := client.FetchObject(v.bucket, key)
	if err != nil && err == riakpbc.ErrObjectNotFound {
		return errors.New("View not found") // TODO: Replace with typed error
	} else if err != nil {
		return err
	}
	reader := bytes.NewReader(rsp.GetContent()[0].GetValue())
	dec := gob.NewDecoder(reader)

	if err := dec.Decode(out); err != nil {
		return err
	}
	return nil
}

func (v *ViewStore) Delete(key string) error {
	client, err := v.connect()
	if err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.DeleteObject(v.bucket, key); err != nil {
		return err
	}

	return nil
}
