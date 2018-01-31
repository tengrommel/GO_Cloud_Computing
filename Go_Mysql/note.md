# Golang Mysql Tutorial
> In order to do this we'll be using https://github.com/go-sql-drive/mysql as our MySQL driver.

Go-SQL-Driver is a lightweight and fast MySQL driver that supports connections over TCP/IPv4,TCP/IPv6,
Unix domain sockets or custom protocols and features automatic handling of broken connections.


## Basic SQL

// perform a db.Query insert 
    insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
    
// if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
// be careful deferring Queries if you are using transactions
    defer insert.Close()
    
## Populating Structs from Results
> Retrieving a set of results from the database is all well and good,but we need to be able to read these 
results or populating existing structs so that we can parse them and modify them easily.

