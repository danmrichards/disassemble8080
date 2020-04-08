package rom

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Load returns the contents of the ROM at the given path as a byte slice.
func Load(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("ROM path cannot be empty")
	}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	rf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open ROM: %w", err)
	}

	rb, err := ioutil.ReadAll(rf)
	if err != nil {
		return nil, fmt.Errorf("could not read ROM: %w", err)
	}

	return rb, nil
}
