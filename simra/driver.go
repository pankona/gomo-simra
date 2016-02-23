package simra

// Driver represents a scene driver.
type Driver interface {
	// Initialize is called to initialize scene.
	Initialize()

	// Drive is called about 60 times per 1 sec.
	// It is the chance to update sprite information like
	// position, appear/disapper, and change scene.
	Drive()
}
