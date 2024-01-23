package wishlist

import (
	"fmt"
	"os"
)

type Item string

type Wishlist struct {
	Storage WishlistStore
}

// The interface defines the common properties expected from the storage backend.
type WishlistStore interface {
	Items() []Item
	//Add will add an item to the wishlist's persistent storage and return a bool indicating success or failure.
	Add(Item)
}

func NewWishList(ws WishlistStore) Wishlist {
	w := Wishlist{Storage: ws}
	return w
}

func (w Wishlist) Items() []Item {
	//ask the store for the items
	return w.Storage.Items()
}

func (w Wishlist) Size() int {
	return len(w.Storage.Items())
}

func (w Wishlist) List() {
	if w.Size() == 0 {
		fmt.Println("no Items in list! To add new Items run 'wishlist foo'")
		os.Exit(0)
	}

	for _, item := range w.Storage.Items() {
		fmt.Println(item)
	}

	os.Exit(0)
}

func (w *Wishlist) Add(item Item) {
	w.Storage.Add(item)

	return
}
