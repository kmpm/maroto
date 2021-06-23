package maroto

import (
	"fmt"
	"sync"

	"github.com/johnfercher/maroto/pkg/props"
)

var (
	modulesMux sync.RWMutex
	modules    = make(map[string]ModuleInfo)
)

type Module interface {
	MarotoModule() ModuleInfo
}

type Options struct {
	Rect props.Rect
	Text props.Text
	Font props.Font

	String string
	Bytes  []byte
}

type Option func(*Options) error

func Rect(left, top, percent float64, center bool) Option {
	return func(o *Options) error {
		o.Rect = props.Rect{Left: left, Top: top, Percent: percent, Center: center}
		return nil
	}
}

func String(s string) Option {
	return func(o *Options) error {
		o.String = s
		return nil
	}
}

type CellContent interface {
	Add(options Options) error
}

type ModuleID string
type ModuleInfo struct {
	ID  ModuleID
	New func() Module
}

func (mi ModuleInfo) String() string { return string(mi.ID) }

func RegisterModule(instance Module) {
	mod := instance.MarotoModule()
	if mod.ID == "" {
		panic("module ID is missing")
	}
	if mod.New == nil {
		panic("missing ModuleInfo.New")
	}
	if val := mod.New(); val == nil {
		panic("ModuleInfo.New must return a non nil instance")
	}
	modulesMux.Lock()
	defer modulesMux.Unlock()

	modules[string(mod.ID)] = mod
}

func Add(id string, options ...Option) error {
	val := modules[id].New().(interface{})

	if val == nil {
		return fmt.Errorf("module can not be nil")
	}
	if prov, ok := val.(CellContent); ok {
		o := Options{}
		for _, cb := range options {
			cb(&o)
		}
		return prov.Add(o)
	}
	return fmt.Errorf("module does not implement CellContent")

}
