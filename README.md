# 🚧 Walley-DB
Hobby binary-based database implementation in Golang

## Usage
Creating a new table
```sh
$ go run main.go CREATE TABLE <table_name> <column_name>:<data_type>=<byte_length>
```
Inserting a record into table
```sh
$ go run main.go INSERT TABLE <table_name> <column_name>=<value> ...
```
Walley-DB files are stored under `$HOME/.walleydb`

^^ table is created at `~/.walleydb/<table_name>`

Each table has it's column configuration saved in `~/.walleydb/<table_name>/<table_name>.conf` and the configuration looks as this:
```
<column_name>:<data_type>=<byte_length>\n
```

### ...so for example
Create new table called `users` with 2 columns:
- id of type int with byte size 8B (int64)
- name of type string with byte size 20B (20 base utf8 encoded chars)
```sh
$ go run main.go CREATE TABLE users id:int=8 name:string=20
```

this saves the whole table under `~/.walleydb/users` and configuration file `~/.walleydb/users/users.conf` will look like this:
```
id:int=8
name:string=20
```

Insert a new record into this table
```sh
$ go run main.go INSERT TABLE users id=123 name=JohnDoe
```

this saves this new user into `~/.walley-db/users/users.wdb` as a fixed length binary record like
```
00 00 00 00 00 31 32 33 (8B of id)
00 00 00 00 00 00 00 00 00 00 00 00 00 4A 6F 68 6E 44 6F 65 (20B of name)
```
^^ note that with table configured to have `8B` for `id` and `20B` for `name`, the whole single record has total length of `28B`.

^^^^ note that currently everything is being written as a `UTF-8` character.

## 👷‍♂️ Under construction
The whole thing is at the very beginnings, I'm trying to learn a lot of things during this time and will update this projects based on my accomplishments. It's not necessarily meant to be a production-grade DB, just an experiment on what I can come up with.
