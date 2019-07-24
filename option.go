package testgrounds

type Option interface {
	Type() string
	Subtype() string
}

type Options []Option

const (
	OPTION_TYPE_STORAGE = "storage"
)

func (o Options) BySubtype(name string) Options {
	var options = make(Options, 0)

	for _, opt := range o {
		if opt.Subtype() == name {
			options = append(options, opt)
		}
	}

	return options
}
