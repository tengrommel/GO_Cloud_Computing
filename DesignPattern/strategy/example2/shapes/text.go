package shapes

import "github.com/tengrommel/GO_Cloud_Computing/DesignPattern/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}