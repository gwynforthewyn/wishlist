package wishlist

import (
	"bufio"
	"fmt"
	"os"
)

type Wishlist struct {
	Items []string
}

func Open(wishlistPath string) (*Wishlist, error) {
	f, err := os.Open(wishlistPath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	w := &Wishlist{}
	input := bufio.NewScanner(f)

	for input.Scan() {
		w.Items = append(w.Items, input.Text())
	}

	return w, nil
}

func (w Wishlist) Size() int {
	return len(w.Items)
}

func (w Wishlist) List() {
	if w.Size() == 0 {
		fmt.Println("no Items in list! To add new Items run 'wishlist foo'")
		os.Exit(0)
	}

	for _, item := range w.Items {
		fmt.Println(item)
	}

	os.Exit(0)
}

func (w *Wishlist) Add(item string) {
	w.Items = append(w.Items, item)

	return
}

func (w Wishlist) Save() {
	f, err := os.Create("wishlist.txt")

	if err != nil {
		fmt.Printf("Failed to open wishlist: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	for _, item := range w.Items {
		_, err = fmt.Fprintln(f, item)
	}

	if err != nil {
		fmt.Printf("Failed to write to wishlist: %v\n", err)
		os.Exit(1)
	}
}
