package logic

import "fca/dal"

var Dao dal.BaseDao

type LogicBase struct {
}

func InitLogic() {

	Dao = dal.BaseDao{}
}
