package window

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
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

// DebugImage saves an image with the specified name, and opens it in a viewer.
func DebugImage(img image.Image, nameFmt string, args ...interface{}) {
	name := fmt.Sprintf(nameFmt, args...)
	pid := fmt.Sprintf("%v", os.Getpid())
	dir := "./debug_img/"
	file := "[" + pid + "] " + name + ".png"

	exists := true

	_, err := os.Stat(dir + file)
	if err != nil {
		exists = false
	}

	os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(dir + file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	f.Sync()

	if exists {
		return
	}

	cmd := exec.Command("xdg-open", dir+file)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
}
