# Go with PostgreSQL

## PostgreSQL with Docker

为了便于测试，可以使用 docker 快速启动一个PostgreSQL数据库。

```
docker run --rm --name my-postgres -e POSTGRES_PASSWORD=yoursecretpassword postgres:17.2-alpine3.20
```

## 驱动包

Go语言支持PostgreSQL数据库驱动程序，包括两个接口：database/sql和pgx。
- pgx 接口更快。许多 PostgreSQL 特定功能（如 LISTEN / NOTIFY 和 COPY ）无法通过 database/sql 界面使用。

在以下情况下，建议使用 pgx 接口：

- 该应用程序仅面向 PostgreSQL。
- 需要使用 PostgreSQL 特定的功能，比如 LISTEN / NOTIFY 和 COPY、service_file等。
- 没有其他需要 database/sql 的库正在使用中。

## 使用 pq

在Go语言中，常用的PostgreSQL数据库驱动包是github.com/lib/pq。

- [pkg.go.dev/github.com/lib/pq](https://pkg.go.dev/github.com/lib/pq)

pq 驱动程序是一个更高级别的接口，它提供了标准 database/sql 接口的实现。

## 使用 pgx

- [github.com/jackc/pgx/v5](https://pkg.go.dev/github.com/jackc/pgx/v5)

pgx 驱动程序是一个更低级别、高性能的接口，它公开了 PostgreSQL 特定的功能，例如 LISTEN / NOTIFY 和 COPY.它还包括一个用于标准 database/sql 接口的适配器。

- [pgx wiki](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx)

## 使用 gorm

此外，gorm也有内置支持的驱动包，可以简化数据库操作。

- [gorm.io/driver/postgres](https://gorm.io/driver/postgres)

安装依赖

首先，安装 gorm 和 postgres 驱动包

```bash
go get -u gorm.io/gorm
```

然后，安装 postgres 驱动包

```bash
go get -u gorm.io/driver/postgres
```

### CRUD 示例

以下是一个简单的CRUD示例：

> 注意:根据你的实际数据库连接信息替换dsn中的参数。运行测试时，确保PostgreSQL数据库服务正在运行，并且数据库和用户已正确配置。

```
// main.go
package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string
}

func main() {
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // 自动迁移模式
    db.AutoMigrate(&User{})

    // 创建
    db.Create(&User{Name: "Alice", Email: "alice@example.com"})

    // 读取
    var user User
    db.First(&user, 1) // 根据整型主键查找
    db.First(&user, "email = ?", "alice@example.com") // 查找email为alice@example.com的用户

    // 更新 - 更新用户的email
    db.Model(&user).Update("Email", "alice@newdomain.com")

    // 删除 - 删除用户
    db.Delete(&user, 1)
}

```

### 测试代码

可以使用testing包来编写测试：

```
// main_test.go
package main

import (
    "testing"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func TestCRUD(t *testing.T) {
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatal(err)
    }

    db.AutoMigrate(&User{})

    // 创建
    user := User{Name: "Bob", Email: "bob@example.com"}
    db.Create(&user)

    // 读取
    var readUser User
    db.First(&readUser, "email = ?", "bob@example.com")
    if readUser.Name != "Bob" {
        t.Errorf("expected name to be Bob, got %s", readUser.Name)
    }

    // 更新
    db.Model(&readUser).Update("Email", "bob@newdomain.com")
    db.First(&readUser, "email = ?", "bob@newdomain.com")
    if readUser.Email != "bob@newdomain.com" {
        t.Errorf("expected email to be bob@newdomain.com, got %s", readUser.Email)
    }

    // 删除
    db.Delete(&readUser)
    var count int64
    db.Model(&User{}).Where("email = ?", "bob@newdomain.com").Count(&count)
    if count != 0 {
        t.Errorf("expected count to be 0, got %d", count)
    }
}

```
