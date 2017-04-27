package users

import (
	"opms/models"

	"github.com/astaxie/beego/orm"
)

type Permissions struct {
	Id         int64 `orm:"pk;column(userid);"`
	Permission string
	Model      string
	Modelc     string
}

func (this *Permissions) TableName() string {
	return models.TableName("permissions")
}
func init() {
	orm.RegisterModel(new(Permissions))
}
