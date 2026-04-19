package craft

// Basicamente sintetizar é processar uma série de elementos e torná-los num material

type SynthesizeOperation struct {
	ID OperationID
	Operation
	Input  Composite // Aqui entram as substâncias necessárias e ficam os resíduos/excedentes (Escória/Gases)
	Output Material  // Aqui sai o produto desejado
}

func (s SynthesizeOperation) GetOperation() Operation { return s.Operation }
func (s SynthesizeOperation) Kind() string            { return "synthesize" }

func (m *Machine) finishSynthesization(s SynthesizeOperation, dt float32) {
	// Lógica:
	// 1. Consome os itens (Composites) do estoque de entrada
	// 2. Cria o novo Material (ex: Barra de Aço)
}

func (m *Machine) calculateSyntheseProduction(s SynthesizeOperation, dt float32) []Material {
	var produced []Material

	return produced
}
