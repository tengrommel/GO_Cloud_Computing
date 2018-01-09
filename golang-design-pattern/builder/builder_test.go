package builder

import "testing"

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123"{
		t.Errorf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder2 := &Builder2{}
	director := NewDirector(builder2)
	director.Construct()
	res := builder2.GetResult()
	if res != 6{
		t.Errorf("Buidler2 fail expect 6 acture %s", res)
	}
}