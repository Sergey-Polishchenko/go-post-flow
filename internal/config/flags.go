package config

import "flag"

type Flags struct {
	InMemory bool
}

func ParseFlags() *Flags {
	flags := &Flags{}
	flag.BoolVar(&flags.InMemory, "inmemory", false, "use inmemory storage")
	flag.Parse()
	return flags
}
