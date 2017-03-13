package dal

import (
	"fca/model"
)

func (dao BaseDao) GetUserByUID(uid int) (user *model.User, err error) {

	/*	obj, err := MDbMap.Get(model.User{}, uid)
		if obj != nil {
			user = obj.(*model.User)

		}
	*/
	return
}
func (dao BaseDao) GetUserByMobile(mobile string) (user *model.User, err error) {
	/*	user = &model.User{}
		err = MDbMap.SelectOne(user, "select * from Users where mobile=?", mobile)
	*/return
}
func (dao BaseDao) AddUser(user model.User) (err error) {
	return
}
