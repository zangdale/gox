package chrome

import (
	"github.com/zserge/lorca"
)

func New(url, dir string, width, height int, customArgs ...string) (lorca.UI, error) {
	return lorca.New(url, dir, width, height, customArgs...)
}
