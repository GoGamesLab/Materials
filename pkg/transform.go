package materials

// Transformar é basicamente mudar a forma de um material em outra forma

type TransformOperation struct {
	ID OperationID
	Operation
	Input  MaterialID // Aqui entra o produto necessário num formato
	Output MaterialID // Aqui sai o produto em outro formato
}

func (s TransformOperation) GetOperation() Operation { return s.Operation }
func (s TransformOperation) Kind() string            { return "transform" }

func (t TransformOperation) Execute(m *Machine, dt float32) {
	m.finishTransformation(t, dt)
}

func (m *Machine) finishTransformation(t TransformOperation, dt float32) {
	//
}
