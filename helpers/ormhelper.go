package helpers

import "github.com/astaxie/beego/orm"

type ormhelper struct {
	Ormer orm.Ormer
}

func (o *ormhelper) getOrm() orm.Ormer {

	if o.Ormer == nil {
		o.Ormer = orm.NewOrm()
	}

	return o.Ormer

}
