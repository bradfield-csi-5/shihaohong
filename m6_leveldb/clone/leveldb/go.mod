module github.com/shihaohong/leveldb_clone/leveldb

go 1.21.1

replace github.com/shihaohong/leveldb_clone/memtable => ../memtable/

require github.com/shihaohong/leveldb_clone/memtable v0.0.0-00010101000000-000000000000
