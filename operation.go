package calc

// Operation is a structure that contains lamda to calculate operation value and
// arguments validation.
type Operation struct {
	lambda     Lambda
	validation Validation
}

var add = Operation{
	lambda:     Add,
	validation: AddValidation,
}

var sub = Operation{
	lambda:     Sub,
	validation: SubValidation,
}

var mul = Operation{
	lambda:     Mul,
	validation: MulValidation,
}

var div = Operation{
	lambda:     Div,
	validation: DivValidation,
}

var sqrt = Operation{
	lambda:     Sqrt,
	validation: SqrtValidation,
}

// Operations supported operations list.
var Operations = map[string]Operation{
	"add":  add,
	"sub":  sub,
	"mul":  mul,
	"div":  div,
	"sqrt": sqrt,
}
