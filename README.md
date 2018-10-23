# sprocbind

Generates go bindings to Microsoft SQL Server (mssql) stored procedures, including Table Valued Parameters (TVP).

It is intended to be used with https://github.com/denisenkom/go-mssqldb once https://github.com/denisenkom/go-mssqldb/pull/413 is merged in.

Until that PR is accepted, you can use https://github.com/elliots/go-mssqldb/tree/tvp

WIP, but in use.

## Usage

`go install github.com/identitii/sprocbind`

`sprocbind --input my-file-with-sprocs.sql --output sprocs.go --go-package db`

## TODO

 - A lot of cleanup.


## Thanks

The authors of the tsql antl4 grammar that this package is built on:

 - Ivan Kochurkin (kvanttt@gmail.com), Positive Technologies.
 - Scott Ure (scott@redstormsoftware.com).
 - Rui Zhang (ruizhang.ccs@gmail.com).
 - Marcus Henriksson (kuseman80@gmail.com).

(TODO: add link to original source of tsql.g4)