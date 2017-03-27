package window

import (
	"fmt"
	"image"
	"os/exec"
	"strings"
	"time"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

// xwin is the type definition for X11 window.
type xwin xwindow.Window

// xconn is the global X-server connection.
var xconn *xgbutil.XUtil

// FindWindow returns a new X11 window based on window name.
func findWindow(winName string) (Window, error) {

	var findErr = fmt.Errorf("Failed to find window by name (name = %v)", winName)

	// Connection has not been established.
	if xconn == nil {

		var err error

		// Create new connection.
		xconn, err = xgbutil.NewConn()

		if err != nil {
			return nil, findErr
		}
	}

	// Get a list of windows.
	wins, err := ewmh.ClientListGet(xconn)
	if err != nil {
		return nil, findErr
	}

	for _, win := range wins {
		name, err := ewmh.WmNameGet(xconn, win)
		if err != nil {
			return nil, findErr
		}

		if strings.Contains(strings.ToLower(name), strings.ToLower(winName)) {
			// Construct window structure from ID.
			return (*xwin)(xwindow.New(xconn, xproto.Window(win))), nil
		}
	}

	return nil, findErr
}

// NewWindow returns a new X11 window based on window ID.
func newWindow(winID int) (Window, error) {

	// Connection has not been established.
	if xconn == nil {

		var err error

		// Create new connection.
		xconn, err = xgbutil.NewConn()

		if err != nil {
			return nil, fmt.Errorf("Failed to attach to window (id = %v)", winID)
		}
	}

	// Construct window structure from ID.
	return (*xwin)(xwindow.New(xconn, xproto.Window(winID))), nil
}

// Resize resizes the X11 window.
func (win *xwin) Resize(width, height int) (err error) {

	err = (*xwindow.Window)(win).WMResize(width, height)
	if err != nil {
		return fmt.Errorf("Failed to resize window (id = %v)", win.Id)
	}

	time.Sleep(time.Millisecond * 500)
	return nil
}

// Image returns a screenshot of the X11 window.
func (win *xwin) Image() (image.Image, error) {

	img, err := xgraphics.NewDrawable(win.X, xproto.Drawable(win.Id))
	if err != nil {
		return nil, fmt.Errorf("Failed to get image from window (id = %v)", win.Id)
	}

	return img, nil
}

// Name returns the window name.
func (win *xwin) Name() (string, error) {

	var nameError = fmt.Errorf("Failed to get window name (id = %v)", win.Id)

	aname := "_NET_WM_NAME"

	nameAtom, err := xproto.InternAtom(win.X.Conn(), true, uint16(len(aname)),
		aname).Reply()
	if err != nil {
		return "", nameError
	}

	reply, err := xproto.GetProperty(win.X.Conn(), false, win.Id, nameAtom.Atom,
		xproto.GetPropertyTypeAny, 0, 0xFFFF).Reply()

	if err != nil {
		return "", nameError
	}

	return string(reply.Value), err
}

// Process gets the window's parrent process name.
func (win *xwin) Process() (string, error) {

	var nameError = fmt.Errorf("Failed to get window name (id = %v)", win.Id)

	aname := "WM_CLASS"

	nameAtom, err := xproto.InternAtom(win.X.Conn(), true, uint16(len(aname)),
		aname).Reply()
	if err != nil {
		return "", nameError
	}

	reply, err := xproto.GetProperty(win.X.Conn(), false, win.Id, nameAtom.Atom,
		xproto.GetPropertyTypeAny, 0, 0xFFFF).Reply()

	if err != nil {
		return "", nameError
	}

	split := strings.Split(string(reply.Value), ".")

	return strings.ToLower(split[0]), err
}

// String returns string representation of window id.
func (win *xwin) String() string {
	return string(int(win.Id))
}

// Send keystroke to window.
// Note: puts window in focus.
func (win *xwin) PressKey(key string) error {
	w := (*xwindow.Window)(win)
	w.Focus()

	cmd := exec.Command("xdotool", "key", key)
	return cmd.Run()
}
