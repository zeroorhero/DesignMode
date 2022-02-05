package abstractfactory

import (
	"testing"
)

func TestAbstract(t *testing.T){
	var fac DaoFactory
	fac = &RGBFactory{}
	fac.CreatDetailDao().SaveDetailDao()
	fac.CreatMainDao().SaveMainDao()
}
