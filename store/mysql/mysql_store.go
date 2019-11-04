package mysql

import (
	"database/sql"
	"errors"
)

type User struct {
	Id   string
	Name string
}

type UserStore struct {
	Db *sql.DB
}

// 新增用户
// 如果插入数据库成功, 则返回id,
// 如果插入数据库失败, 则返回错误.
func (userStore *UserStore) Create(user *User) (int64, error) {
	query := ""
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 更新用户
// 如果更新数据行成功, 则返回受影响的行数,
// 如果更新数据行失败, 则返回错误.
func (userStore *UserStore) Update(user *User) (int64, error) {
	query := ""
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 查询指定用户
func (userStore *UserStore) User(id string) (*User, error) {
	var user User
	query := "select * from t_article t where t.id = $1"
	row := userStore.Db.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

// 查询用户列表
func (userStore *UserStore) Users() ([]*User, error) {
	query := "select * from t_article"
	_, err := userStore.Db.Query(query)
	if err != nil {
		return nil, err
	}
	err = errors.New("未实现, 如何解析成列表?")
	return nil, err
}

// 删除所有用户
func (userStore *UserStore) RemoveAll() (rowsAffected int64, err error) {
	query := "truncat table t_article"
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 删除指定用户
func (userStore *UserStore) Remove(id string) (rowsAffected int64, err error) {
	query := "delete from t_article t where t.id = $1"
	result, err := userStore.Db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
