package slugs

import (
	"reflect"
	"testing"

	"github.com/remind101/empire/images"
)

func TestRepository(t *testing.T) {
	r := newRepository()

	if got, want := len(r.slugs), 0; got != want {
		t.Fatal("Expected no slugs")
	}

	image := &images.Image{Repo: "remind101/r101-api", ID: "1234"}
	if slug, err := r.Create(&Slug{Image: image}); err == nil {
		expected := &Slug{
			ID:    "1",
			Image: image,
		}
		if got, want := slug, expected; !reflect.DeepEqual(got, want) {
			t.Fatalf("Create => %q; want %q")
		}
	} else {
		t.Fatal(err)
	}

	if got, want := len(r.slugs), 1; got != want {
		t.Fatal("Slugs count %d; want %d", got, want)
	}

	if slug, err := r.FindByImage(&images.Image{Repo: "remind101/r101-api", ID: "1234"}); err == nil {
		if slug == nil {
			t.Fatal("Expected a slug to be returned")
		}
	} else {
		t.Fatal(err)
	}
}
