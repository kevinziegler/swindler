package adapter

type Adapter interface {
	Show(string)
}

/*
  {
    noop: Noop.new
    dowop: dowop.new
  }

  class Noop
    def show(thing)...
  end
*/
