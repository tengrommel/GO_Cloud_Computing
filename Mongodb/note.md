## mongodb

### Nosql

- Collections instead of Tables
- Database is a set of collections
- Collection is a container for document
- Documents are simple "structs"

### MongoDB

- Mongo shell uses JavaScript-like noation with JSON
- Documents are JSON structs
- _id is the unique key of every single document
- Document can be embedded or referenced by _id(ObjectId type)

### MongdDB: useful comments

- show dbs
- show collections
- use <dbname>
- db <collectionname> 
    > .find() returns a set(collection)<br>
    .find().toArray()<br>
    .findOne()<br>
    .count()<br>
    .insert()<br>
    .save()<br>
    .remove()<br>
    .drop()<br>
    

### Characteristics

- No transactions
- Atomic updates/transactions only on the level of a single document
