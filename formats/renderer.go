package formats

import "github.com/maxlinc/tabler"

type Renderer interface {
	Render(tabler.Table) error
}
