package vehicle

type VehicleInterface interface {
	MoveFoward(speed int) (bool, error)
}
