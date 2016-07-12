# Go ORM Benchmark By PostgreSQL

### Environment

* go version go1.6 linux/amd64

### PostgreSQL

* PostgreSQL 9.5 for Linux on x86_64

### ORMs

All package run in no-cache mode.

* [dbr](https://github.com/gocraft/dbr) (in preparation)
* [genmai](https://github.com/naoina/genmai) (in preparation)
* [gorm](https://github.com/jinzhu/gorm)
* [gorp](https://github.com/go-gorp/gorp) (in preparation)
* [pg](https://github.com/go-pg/pg)
* [beego/orm](https://github.com/astaxie/beego/tree/master/orm)
* [sqlx](https://github.com/jmoiron/sqlx) (in preparation)
* [xorm](https://github.com/xormplus/xorm)
	
### Init

```go
# build
go install
# table create, test data insert.
go-orm-benchmark-by-pg
```

### Run

```go
# change dir orms.
cd orms
# All
go test -bench . -benchmem
# Single
go test -bench PqRead -benchmem

go test -bench . -benchmem | prettybench
```

### Reports

* ベンチマーク名
* 総実行時間（少ないほど良い）
* １回の実行にかかった時間（少ないほど良い）
* 実行ごとに割り当てられたメモリのサイズ（少ないほど良い）
* １回の実行でメモリアロケーション（メモリ割り当て/配分）が行われた回数（少ないほど良い）
