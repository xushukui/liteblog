package models

//User 用户表
type User struct {
	Model
	Name   string `gorm:"unique_index" json:"name"`
	Email  string `gorm:"unique_index" json:"email"`
	Pwd    string `json:"-"`
	Avatar string `json:"avatar"`
	Role   int    `gorm:"default:0" json:"role"` //0代表管理员, 1代表普通用户
	Editor string `json:"editor"`
}

//QueryUserByEmailAndPassword 根据邮箱和密码查询用户
func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

//QueryUserByName 根据昵称查询用户
func (db *DB) QueryUserByName(name string) (user User, err error) {
	return user, db.db.Where("name = ?", name).Take(&user).Error
}

//QueryUserByEmail 根据邮箱查询用户
func (db *DB) QueryUserByEmail(email string) (user User, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

//UpdateUserEditor 更改用户信息
func (db *DB) UpdateUserEditor(editor string) (err error) {
	return db.db.Model(&User{}).Update("editor", editor).Error
}

//SaveUser 保存新注册的用户
func SaveUser(user *User) error {
	return db.Create(user).Error
}
