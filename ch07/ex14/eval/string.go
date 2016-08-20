package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	s := c.fn + "("
	for i, arg := range c.args {
		if i > 0 {
			s += ","
		}
		s += arg.String()
	}
	s += ")"
	return s
}

func (m min) String() string {
	s := "min("
	for i, arg := range m.args {
		if i > 0 {
			s += ","
		}
		s += arg.String()
	}
	s += ")"
	return s
}
