package image

import (
	"log"

	"github.com/johnfercher/maroto"
)

type Provider struct{}

func init() {
	maroto.RegisterModule(Provider{})
}

func (p Provider) MarotoModule() maroto.ModuleInfo {
	return maroto.ModuleInfo{
		ID:  "image",
		New: func() maroto.Module { return &Provider{} },
	}
}

func (p *Provider) Add(options maroto.Options) error {
	log.Print("asdfasdf")
	return nil
}

var (
	_ maroto.CellContent = (*Provider)(nil)
)
