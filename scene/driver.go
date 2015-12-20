package scene

type Driver interface {
	Initialize(func(nextScene Driver))
	Drive()
}
