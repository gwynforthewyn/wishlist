package wishlist

import (
	"bufio"
	"fmt"
	"os"
)

type FileStore struct {
	WishlistPath  string
	wishlistItems []Item
}

func OpenFileStore(wishlistPath string) (*FileStore, error) {

	store := FileStore{
		WishlistPath: wishlistPath,
	}

	f, err := os.Open(store.WishlistPath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		store.wishlistItems = append(store.wishlistItems, Item(input.Text()))
	}

	return &store, nil
}

func (f FileStore) Add(item Item) {
	f.wishlistItems = append(f.wishlistItems, item)
	f.save()
}

func (f FileStore) Items() []Item {
	return f.wishlistItems
}

func (f FileStore) save() {
	file, err := os.Create(f.WishlistPath)

	if err != nil {
		fmt.Printf("Failed to open wishlist: %v\n", err)
		os.Exit(1)
	}

	//return, sync, defer, close
	defer file.Close()

	for _, item := range f.Items() {
		_, err = fmt.Fprintln(file, item)
	}

	if err != nil {
		fmt.Printf("Failed to write to wishlist: %v\n", err)
		os.Exit(1)
	}

	//explicitly sync to validate no problems writing to disk
	err = file.Sync()

	if err != nil {
		fmt.Printf("Wishlist failed to write to disk!: %v\n", err)
	}
}
