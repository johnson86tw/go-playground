package builder

import "testing"

func TestBuildCar(t *testing.T) {
	builder := NewBuilder()
	builder.Paint(RedColor).Wheels(SteelWheels).TopSpeed(MPH)
	familyCar := builder.Build()
	familyCar.Drive()

	sportsCar := builder.Wheels(SportsWheels).TopSpeed(KPH).Build()
	sportsCar.Drive()

}
