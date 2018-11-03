package adapter

type Adapter interface {
	Add(string)
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
