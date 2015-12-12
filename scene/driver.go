package scene

type Driver interface {
	Initialize(func(d Driver))
	Drive()
}
