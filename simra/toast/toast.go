package toast

type Toaster interface {
	Show(text string) error
}
