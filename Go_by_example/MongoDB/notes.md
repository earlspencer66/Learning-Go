### Learning MongoDB
1. MongoDB locally 
<https://www.mongodb.com/try/download/community>

    <https://fastdl.mongodb.org/windows/mongodb-windows-x86_64-7.0.8-signed.msi>

2. MongoDB Shell
<https://www.mongodb.com/try/download/shell>

3. launch mongo shell
```cmd
mongosh
```

#### Collections and Documents
* Data stored in *collections*
* *Collection* contains *documents*
```bson
{
    "Name": "",
    "Author": "",
    "_id": ObjectId("2326236"),
}
```

#### MongoDB Compasss

#### Using the MongoDB Shell
```cmd
mongosh

test> show.dbs
test> use local
local> 
```

#### Adding New Documents
```cmd
bookstore> db.books
bookstore.books
bookstore> db.books.insertOne({ json file})
```
Create a collection and add documents
```cmd
bookstore> db.authors.insertOne({json file})
```

Insert multiple documents at once
```cmd
bookstore> db.books.insertMany([{json file}, {json file}])
```
#### Finding Documents
```cmd
bookstore> db.books.find()
//prints out the first 20 books
bookstore> it
//iterates over the next 20 books

bookstore>db.books.find({author: "Ellen G White})
//find by filter

bookstore> db.books.find({author: "Ellen G White"}, {author: 1, year:1})
//ony return specific properties of the books

bookstore> db.books.find({}, {author:1, series:1})
//returns all the books but 
//specific properties
```

#### Sorting and Limiting Data
```cmd
bookstore> db.books.find().count()
//chaining

bookstore> db.books.find().limit(3)

bookstore> db.books.find().sort({author: 1})
//ascending order

bookstore> db.books.find().sort({author:-1})
//descending order
```
