package builder

import "testing"

func TestBuilder_Build(t *testing.T) {
	house := &House{}
	builderHouse := &Builder{house:house}
	h :=builderHouse.Color("red").Area(60.45).Material("wood").Build()
	t.Logf("%+v\n",h)
}
