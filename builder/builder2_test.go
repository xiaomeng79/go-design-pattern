package builder

import "testing"

func TestDoorBuilder_Build(t *testing.T) {
	door := DoorBuilder{}.Build(new(Door),new(Lock))
	t.Logf("价格:%f\n",door.GetCost())
	t.Logf("包含:%s\n",door.ShowItems())

	door = DoorBuilder{}.Build(new(Door),new(Lock),new(Handle))
	t.Logf("价格:%f\n",door.GetCost())
	t.Logf("包含:%s\n",door.ShowItems())
}
