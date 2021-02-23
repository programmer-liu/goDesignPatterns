package OOP

type Operator interface {//多态
	 GetResult()(result float64)
}

type Add struct {//封装
	NumberA float64
	NumberB float64
}

func (add *Add)GetResult()(result float64){
	return add.NumberA+add.NumberB
}

type Sub struct {
	NumberA float64
	NumberB float64
}
func (sub *Sub)GetResult()(result float64){
	return sub.NumberA-sub.NumberB
}

type Multi struct {
	NumberA float64
	NumberB float64
}
func (multi *Multi)GetResult()(result float64){
	return multi.NumberA*multi.NumberB
}

type Divis struct {
	NumberA float64
	NumberB float64
}
func (divis *Divis)GetResult()(result float64){
	return divis.NumberA/divis.NumberB

}

