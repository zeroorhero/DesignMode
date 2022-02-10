package stratege

import "fmt"

type conte struct {
	name string
	money int
}

type stratege1 interface {
	Pay1(con *conte)
}

type cssh struct {

}
func (c *cssh)Pay1(con *conte){
	fmt.Println("pay %s %d via cash", con.name, con.money)
}

type Payment struct {
	str stratege1
	con conte
}
func (p *Payment) Pay(){
	p.str.Pay1(p.con)
}