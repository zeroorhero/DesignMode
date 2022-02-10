package stratege

import "fmt"

type context struct {
	Name, CardID string
	Money        int
}

type stratege interface {
	Pay(con *context)
}

type Cash struct {}
func (c *Cash) Pay(con *context){
	fmt.Println("via cash to pay")
}

type Bank struct {}
func (c *Bank) Pay(con *context){
	fmt.Println("via Bank to pay")
}

type PayMent struct {
	con *context
	stra stratege
}

func NewPayMent(Name, CardID string, money int, stra stratege) *PayMent{
	return &PayMent{
		con:  &context{
			Name:   Name,
			CardID: CardID,
			Money:  money,
		},
		stra: stra,
	}
}
func (p *PayMent)Pay(){
	p.stra.Pay(p.con)
}