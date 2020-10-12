//Create and switch database
use bookdb

//Create collection
db.createCollection("book")

//Find all book
db.book.find()

//Remove all book
db.book.remove({})

