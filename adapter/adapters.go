package adapter

var Adapters = map[string]Adapter {
	"noop": Noop{},
	"dowop": Dowop{},
}
