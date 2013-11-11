package riak_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"persistence/riak"
	"reflect"
	"testing"
)

type testView struct {
	TestField string
}

type testView2 struct {
	InnerView testView
}

func TestStoreView(t *testing.T) {
	t.Log("When a view is stored under a key")

	key := uuid.New()
	storedView := testView{TestField: "hello"}

	if err := Store(key, storedView); err != nil {
		t.Fatal(err)
	}
	defer Delete(key)

	receivedView := testView{}
	if err := Get(key, &receivedView); err != nil {
		t.Fatal(err)
	}

	t.Log("It should match the value retrieved using the key")
	if reflect.DeepEqual(storedView, receivedView) == false {
		t.Errorf("\tExpected %v but was %v", storedView, receivedView)
	}
}

func TestStoreAnotherViewType(t *testing.T) {
	t.Log("When a view of another type is stored under a key")

	key := uuid.New()
	storedView := testView2{InnerView: testView{TestField: "Madness, as you know, is a lot like gravity, all it takes is a little push."}}

	if err := Store(key, storedView); err != nil {
		t.Fatal(err)
	}
	defer Delete(key)

	receivedView := testView2{}
	if err := Get(key, &receivedView); err != nil {
		t.Fatal(err)
	}

	t.Log("It should match the value retrieved using the key")
	if reflect.DeepEqual(storedView, receivedView) == false {
		t.Errorf("\tExpected %v but was %v", storedView, receivedView)
	}
}

func TestStoreDifferentViews(t *testing.T) {
	t.Log("When multiple views are stored under different keys")

	key1 := uuid.New()
	storedView1 := testView{TestField: "hello"}
	key2 := uuid.New()
	storedView2 := testView{TestField: "hello again"}

	if err := Store(key1, storedView1); err != nil {
		t.Fatal(err)
	}
	defer Delete(key1)

	if err := Store(key2, storedView2); err != nil {
		t.Fatal(err)
	}
	defer Delete(key2)

	t.Log("\tRetrieve views")
	receivedView1 := testView{}
	if err := Get(key1, &receivedView1); err != nil {
		t.Fatal(err)
	}
	receivedView2 := testView{}
	if err := Get(key2, &receivedView2); err != nil {
		t.Fatal(err)
	}

	t.Log("The retrieved views should match the stored views")
	if reflect.DeepEqual(storedView1, receivedView1) == false {
		t.Errorf("\tExpected %v but was %v", storedView1, receivedView1)
	}
	if reflect.DeepEqual(storedView2, receivedView2) == false {
		t.Errorf("\tExpected %v but was %v", storedView2, receivedView2)
	}
}

func TestStoreDifferentViewValuesUnderSameKey(t *testing.T) {
	t.Log("When different views are stored under the same key")

	key := uuid.New()
	storedView1 := testView{TestField: "hello"}
	storedView2 := testView{TestField: "hello again"}

	if err := Store(key, storedView1); err != nil {
		t.Fatal(err)
	}
	defer Delete(key)
	if err := Store(key, storedView2); err != nil {
		t.Fatal(err)
	}

	receivedView := testView{}
	if err := Get(key, &receivedView); err != nil {
		t.Fatal(err)
	}

	t.Log("It should retrieve the last stored view")
	if reflect.DeepEqual(storedView2, receivedView) == false {
		t.Errorf("\tExpected %v but was %v", storedView2, receivedView)
	}
}

func TestGetNoObject(t *testing.T) {
	t.Log("When no view is stored under a key")

	receivedView := testView{}
	err := Get(uuid.New(), &receivedView)

	t.Log("It should return an error saying the view doesn't exist")
	if err == nil || err.Error() != "View not found" {
		t.Errorf("\tExpected \"%v\" but was %v", "View not found", err)
	}
}

func Store(key string, in interface{}) error {
	return riak.NewViewStore(riakCluster, riakClientId).Store(key, in)
}

func Get(key string, out interface{}) error {
	return riak.NewViewStore(riakCluster, riakClientId).Get(key, out)
}

func Delete(key string) error {
	return riak.NewViewStore(riakCluster, riakClientId).Delete(key)
}
