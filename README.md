# Walley-DB
Hobby binary-based database implementation in Golang

## Usage
```sh
$ go run main.go CREATE TABLE <table_name> <column_name>:<data_type>=<byte_length>
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
