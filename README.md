wishlist - listing the items
wishlist <foo> - adds foo to the list
wishlist -d <foo> - deletes foo from the list

multiple lists!
there's state stored somewhere!

# no items
$ wishlist
>>> no items in list! To add new items run 'wishlist foo'
$

# add an item
$ wishlist https://bitfieldconsulting.com/books/crypto
>>> https://bitfieldconsulting.com/books/crypto added to wishlist

# add same item twice
$ wishlist https://bitfieldconsulting.com/books/crypto
>>> https://bitfieldconsulting.com/books/crypto added to wishlist
$ wishlist https://bitfieldconsulting.com/books/crypto
>>> $ https://bitfieldconsulting.com/books/crypto added to your wishlist  

# with 2 items, that are the same
$ wishlist
>>> https://bitfieldconsulting.com/books/crypto
>>> https://bitfieldconsulting.com/books/crypto

# delete empty list
$ wishlist -d foo

# delete item that doesn't exist
