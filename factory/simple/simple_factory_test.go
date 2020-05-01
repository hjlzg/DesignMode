package simple

import "testing"

func TestSimple(t *testing.T){
	//创建工厂
	var girlFactory = new(GirlFactory) //传递姑娘类型
	girl:=girlFactory.GreateGirl("fat")
	if girl!=nil{
		girl.weight()
	}

	girl=girlFactory.GreateGirl("thin")
	if girl!=nil{
		girl.weight()
	}
}