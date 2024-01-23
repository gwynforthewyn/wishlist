package wishlist_test

import (
	"github.com/gwynforthewyn/wishlist/wishlist"
	"slices"
	"testing"
)

func TestOpenWishlist(t *testing.T) {
	wishlistFixture := "testdata/wishlistOfTwoItems.txt"
	expected := []wishlist.Item{"testItem1", "testItem2"}

	wl, err := wishlist.OpenFileStore(wishlistFixture)

	if err != nil {
		t.Fatalf("Could not open the test %v", err)
	}

	//TODO: Homework: Items must become a method on the wishlist, so the sql store can be the same interface
	if !slices.Equal(wl.Items(), expected) {
		t.Fatalf("Wishlist of two items mismatch. Wanted %s, got %s\n", expected, wl.Items())
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
