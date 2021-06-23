package text

import (
	"fmt"

	"github.com/johnfercher/maroto"
)

type Provider struct{}

func init() {
	maroto.RegisterModule(Provider{})
}

func (p Provider) MarotoModule() maroto.ModuleInfo {
	return maroto.ModuleInfo{
		ID:  "text",
		New: func() maroto.Module { return &Provider{} },
	}
}

func (p *Provider) Add(options maroto.Options) error {
	fmt.Println("text", options.String)
	return nil
}

var (
	_ maroto.CellContent = (*Provider)(nil)
)
