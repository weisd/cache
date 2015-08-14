package cache

import (
	"testing"
)

func Test_TagCache(t *testing.T) {

	c, err := New(Options{Adapter: "memory"})
	if err != nil {
		t.Fatal(err)
	}

	err = c.Put("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}

	res := c.Get("da")

	if res != "weisd" {
		t.Fatal("sdf")
	}

	t.Log("ok")
	t.Log("test", res)

	err = c.Tags([]string{"dd"}).Put("da", "weisd", 300)
	if err != nil {
		t.Fatal(err)
	}
	res = c.Tags([]string{"dd"}).Get("da")

	if res != "weisd" {
		t.Fatal("not weisd")
	}

	t.Log("ok")
	t.Log("dd", res)

	err = c.Tags([]string{"aa"}).Put("aa", "aaa", 300)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Tags([]string{"aa"}).Put("bb", "bbb", 300)
	if err != nil {
		t.Fatal(err)
	}

	res = c.Tags([]string{"aa"}).Get("aa")

	if res != "aaa" {
		t.Fatal("not aaa")
	}

	t.Log("ok")
	t.Log("aa", res)

	err = c.Tags([]string{"aa"}).Flush()
	if err != nil {
		t.Fatal(err)
	}

	res = c.Tags([]string{"aa"}).Get("aa")
	if res != "" {
		t.Fatal("flush faield")
	}

	res = c.Tags([]string{"aa"}).Get("bb")
	if res != "" {
		t.Fatal("flush faield")
	}

	res = c.Tags([]string{"dd"}).Get("da")
	if res != "weisd" {
		t.Fatal("not weisd")
	}

	t.Log("ok")

}
