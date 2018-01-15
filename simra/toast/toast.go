package toast

// Toaster is an interface to handle toast
type Toaster interface {
	// Show shows toast
	Show(text string) error
}
