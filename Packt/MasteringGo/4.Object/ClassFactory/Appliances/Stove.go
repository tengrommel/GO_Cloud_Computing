package Appliances

type Stove struct {
	typeName string
}

func (sv *Stove)Start()  {
	sv.typeName = " Stove "
}

func (sv *Stove)GetPurpose() string {
	return "I am a " + sv.typeName + " I cook food!!"
}