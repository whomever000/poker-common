package window

import (
	"image"
)

var target Window

// Window represents an OS window.
type Window interface {

	// Resize resizes the window. The size is measured excluding window
	// decorations.
	Resize(width, height int) error

	// Image captures the framebuffer of the window.
	Image() (image.Image, error)

	// Name gets the window name.
	Name() (string, error)

	// Process gets the window's parrent process name.
	Process() (string, error)

	// String returns string representation of window id.
	String() string
}

// Attach attaches to a window based on its name.
func Attach(winName string) (Window, error) {
	var err error
	target, err = findWindow(winName)
	return target, err
}

// AttachByID attaches to a window based on its ID.
func AttachByID(winID int) (Window, error) {
	var err error
	target, err = newWindow(winID)
	return target, err
}

// Get returns the target window.
func Get() Window {
	return target
}
