package blackParams

type (
	RequestWithGlobal struct {
		name string
	}

	RequestWithLocal struct {
		r      RequestWithGlobal
		global []int
	}

	Request struct {
		r     RequestWithLocal
		local []int
	}
)

func NewRequest(name string) RequestWithGlobal {
	return RequestWithGlobal{name: name}
}

func (r RequestWithGlobal) WithGlobal(workers ...int) RequestWithLocal {
	return RequestWithLocal{
		r:      r,
		global: workers,
	}
}

func (r RequestWithGlobal) WithoutGlobal() RequestWithLocal {
	return RequestWithLocal{
		r: r,
	}
}

func (r RequestWithLocal) WithLocal(workers ...int) Request {
	return Request{
		r:     r,
		local: workers,
	}
}

func (r RequestWithLocal) WithoutLocal() Request {
	return Request{
		r: r,
	}
}
