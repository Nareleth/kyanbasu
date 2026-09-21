package kyanbasu

import (
    "os"
    "golang.org/x/term"
)


/*
    To be replaced in the future, I just don't know how to do this part myself.
    x/term is utilized to expand the capabilities of the terminal.
*/


/* Types*/

// Raw
type Raw struct {
    fd      int
    state   *term.State
}

/* Functions */


// Get the size of the terminal
func GetSize() (int, int) {
    width, height, _ := term.GetSize(int(os.Stdout.Fd()))
    return width, height
}


// Set Terminal to Raw Mode
func RawMode(f *os.File) (*Raw, error) {
    fd := int(f.Fd())

    state, err := term.MakeRaw(fd)

    if err != nil {
        return nil, err
    }
    return &Raw{fd: fd, state: state}, nil
}

// Return terminal to normal
func (r *Raw) RawClean() error {
    return term.Restore(r.fd, r.state)
}

