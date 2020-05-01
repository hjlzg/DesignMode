package simple

import "fmt"

type Girl interface{
	weight()
}

//胖女孩
type FatGirl struct {
}

type GirlFactory struct {
}

//瘦女孩
type ThinGirl struct{
}

func (FatGirl) weight(){
	fmt.Print("60kg")
}

func (ThinGirl) weight(){
	fmt.Print("50kg")
}

func (*GirlFactory) GreateGirl(like string) Girl{
	if like=="fat"{
		return &FatGirl{}
	}else if like=="thin"{
		return &ThinGirl{}
	}
	return nil
}