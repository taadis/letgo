package user

import (
	"log"

	"github.com/taadis/letgo/store"
)

// 新增用户
// 如果插入数据库成功, 则返回id,
// 如果插入数据库失败, 则返回错误.
/*
func (userStore *UserStore) Create(user *store.User) (int64, error) {
	query := ""
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
*/

// 更新用户
// 如果更新数据行成功, 则返回受影响的行数,
// 如果更新数据行失败, 则返回错误.
/*
func (userStore *UserStore) Update(user *store.User) (int64, error) {
	query := ""
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
*/

// Find
func Find(id string) (User, error) {
	query := "select id, name from system_user t where t.id = ?"
	row := store.Db.QueryRow(query, id)
	var user User
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		log.Fatalln("row.Scan error", err.Error())
		return user, err
	}
	return user, nil
}

// 查询用户列表
func Users() ([]*User, error) {
	query := "select id, name, password, avatar from system_user"
	rows, err := store.Db.Query(query)
	if err != nil {
		log.Println("users query error", err.Error())
		return nil, err
	}
	defer rows.Close()

	// for
	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Avatar)
		if err != nil {
			log.Fatalln("rows.Scan error", err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// 删除所有用户
/*
func (userStore *UserStore) RemoveAll() (rowsAffected int64, err error) {
	query := "truncat table t_article"
	result, err := userStore.Db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
*/

// Remove
func Remove(id string) (int64, error) {
	query := "delete from system_user where t.id = $1"
	result, err := store.Db.Exec(query, id)
	if err != nil {
		log.Fatalln("db exec error", err.Error())
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("result rowsAffected error", err.Error())
		return 0, err
	}
	return rowsAffected, nil
}
