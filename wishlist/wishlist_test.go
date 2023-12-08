package wishlist_test

import (
	"github.com/gwynforthewyn/wishlist/wishlist"
	"slices"
	"testing"
)

func TestOpenWishlist(t *testing.T) {
	wishlistFixture := "testdata/wishlistOfTwoItems.txt"
	expected := []string{"testItem1", "testItem2"}

	wl, err := wishlist.Open(wishlistFixture)

	if err != nil {
		t.Fatalf("Could not open the test %v", err)
	}

	if !slices.Equal(wl.Items, expected) {
		t.Fatalf("Wishlist of two items mismatch. Wanted %s, got %s\n", expected, wl.Items)
	}
}

//func TestOpenNonEmptylist(t *testing.T) {
//	wishlistFixture := "testdata/wishlist.txt"
//	_, err := wishlist.Open(wishlistFixture)
//
//	if err != nil {
//		t.Fatalf("Could not open the test %v", err)
//	}
//}
