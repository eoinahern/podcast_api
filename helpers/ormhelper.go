package helpers

import "github.com/astaxie/beego/orm"

type Ormhelper struct {
	Ormer orm.Ormer
}

func (o *Ormhelper) GetOrm() orm.Ormer {

	if o.Ormer == nil {
		o.Ormer = orm.NewOrm()
		o.Ormer.Using("default")
	}

	return o.Ormer
}
