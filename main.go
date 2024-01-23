package main

import (
	"github.com/gwynforthewyn/wishlist/wishlist"
	"os"
)

func main() {
	defaultWishlistPath := "wishlist.txt"
	//defaultDatabasePath := "wishlist.sqlite"

	//if <some argument> {
	//	wishlistStorage = wishlist.OpenDBStore(defaultDatabasePath)
	//}

	wishlistStorage, err := wishlist.OpenFileStore(defaultWishlistPath)

	if err != nil {
		//There is some error, and the error is not that wishlist.txt doesn't exist.
		//This is triggered if the command _has_ been run before, but wishlist.txt
		//cannot be read for some reason.
		//If the wishlist command has never been run before, then we will have an error
		//but the error is that wishlist.txt doesn't exist, which we can ignore.
		if !os.IsNotExist(err) {
			panic(err)
		}

		//There is an error, the error is the file doesn't exist! We're running for the first
		//time, so initialise the list.
	}

	list := wishlist.NewWishList(wishlistStorage)

	programArgs := os.Args[1:]
	// wishlist -- when list is empty, len(os.args) is 1, wishlist store has things
	// wishlist -- when list is empty, len(os.args) is 1, wishlist store is empty
	shouldWeList := len(programArgs) == 0

	if shouldWeList {
		list.List()
	} else {
		list.Add(wishlist.Item(os.Args[1]))
	}
}
