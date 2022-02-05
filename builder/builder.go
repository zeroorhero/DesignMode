package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}
func (d *Director)construct(){
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
func NewDirector(builder Builder) *Director{
	return &Director{
		builder: builder,
	}
}

type type1 struct {
	res string
}
func (t *type1)Part1(){
	t.res+="1"
}
func (t *type1)Part2(){
	t.res+="2"
}
func (t *type1)Part3(){
	t.res+="3"
}
func (t *type1)GetRes()string{
	return t.res
}

type type2 struct {
	res int
}
func (t *type2)Part1(){
	t.res+=1
}
func (t *type2)Part2(){
	t.res+=2
}
func (t *type2)Part3(){
	t.res+=3
}
func (t *type2)GetRes()int{
	return t.res
}