# TODO


## A nova idéia da Redução

Materiais são reduzidos à substâncias que são reduzidas à elementos sem perdas, usando puramente as composições
A perda é determinada pelo processo
Existe perda por elemento, que determina a perda por substância e que por sua vez determina a perda por material!

Por exemplo:
```go
Material{
    ID:   CoalID,
    Name: "Coal",
    Composites: []Composite{
        {PureCarbonID, 80},
        {PureSulfurID, 5},
        {MethaneID, 15},
    },
}

Substance{
    ID:           PureCarbonID,
    Name:         "Carbon",
    Composition:  []ChemicalBond{
        {Element: CarbonID, Amount: 1},
    },
    MeltingPoint: -78.5,
    BoilingPoint: -78.5,
}

Substance{
    ID:           PureSulfurID,
    Name:         "Sulfur",
    Composition:  []ChemicalBond{
        {Element: SulfurID, Amount: 1},
    },
    MeltingPoint: 115.2,
    BoilingPoint: 444.6,
}

Substance{
    ID:   MethaneID,
    Name: "Methane",
    Composition: []ChemicalBond{
        {Element: CarbonID, Amount: 1},
        {Element: HydrogenID, Amount: 4},
    },
    MeltingPoint: -182.5,
    BoilingPoint: -161.5,
}
```

Executando:
```go
func (m *Material) Reduce(quantity float32) map[SubstanceID]float32 {
	r := make(map[SubstanceID]float32)
	for _, c := range m.Composites {
		r[c.substance] = (c.percentual / 100) * quantity
	}

	return r
}

Coal.Reduce(1)
```

Resulta em:
```go
PureCarbonID = 0.8
PureSulfurID = 0.05
MethaneID = 0.15
```

Em seguida executando:
```go
func (m *Substance) Reduce(quantity float32) map[ElementID]float32 {
	r := make(map[ElementID]float32)
	for _, c := range m.Composition {
		r[c.Element] = float32(c.Amount) * quantity
	}

	return r
}

PureCarbon.Reduce(0.8)
PureSulfur.Reduce(0.05)
Methane.Reduce(0.15)
```

Resulta em:
```go
PureCarbonID = 0.8
PureSulfurID = 0.05
MethaneID = 0.15
```

Para a seguinte operação de queima de carvão (está como refino -- decomposição):
```go
RefineOperation{
    Operation: Operation{
        ID:               CoalCombustionID,
        Name:             "Combustão de Carvão",
        RequiredTemp:     600.0,
        ActivationTemp:   300.0,
        Duration:         10.0,
        EnergyCost:       0,
        BaseEnergyChange: 150.0,
        Dissipation:      3.5,
    },
    Input: CoalID,
    Output: []Composite{
        {PureCarbonID, 5},
        {PureSulfurID, 5},
        {MethaneID, 15},
    },
}
```

Temos os seguintes resíduos (em percentuais)
```go
Output: []Composite{
    {PureCarbonID, 5},
    {PureSulfurID, 5},
    {MethaneID, 15},
},
```

Que aplicadas aos resultados anteriores:
```go
PureCarbonID = 0.8 * 5 / 100 = 0.4
PureSulfurID = 0.05 * 5 / 100 = 0.025
MethaneID = 0.15 * 15 / 100 = 0.15
```

Este material permanece no inventário da máquina, juntamente com o combustível que sobrar


## Temperatura acumulada

A temperatura da máquina sempre se acumula com o passar do tempo. Talvez devesse ser uma curva! Fiz uma função para calcular a temperatura com base na temperatura atual e um fator de ganho que depende do material e condições de queima!


## Duração do processo X material disponível

O sistema está com uma dualidade para determinar o final do processo. Foi idealizado que o material de entrada duraria um período para alimentar o processo, porém calculamos o consumo do material a cada tick e portanto o tempo de duração não está sendo considerado corretamente!

Atualizando uma varíavel `InternalFuel` da máquina, porém como passamos a usar o inventário, InternalFuel perdeu a função, ou deveria ser substituído por Machine.inventory[Machine.Input], porém isso é definido por processo e não por máquina!


## Execução do processo (DONE)

Ao executar o processo, a máquina está consumindo todo o material na primeira iteração!

Em `func (m *Machine) executeStep(step MachineOp, dt float32)` de [crafting.go](./pkg/craft/crafting.go) chamamos as funções de finalização de cada máquina, mas por exemplo `func (m *Machine) finishDistillation(d Refine, dt float32)` em [refine.go](./pkg/craft/refine.go), tem um loop `for m.inventory.Materials[d.Input] > 0` que vai consumir tudo na primeira chamada!

Se mudar para `if m.inventory.Materials[d.Input] > 0` deve resolver!

Porém também tivemos rever as funções que adicionam material no inventário, pois elas estão considerando unidades e usam funções de produção que por sua vez consideram percentuais!


## Conversão de unidades (DONE)

Durante o processamento de materiais por uma máquina, não estamos considerando corretamente a conversão de unidades

Em [refine.go](./pkg/craft/refine.go) por exemplo, durante o refinamento subtraimos uma unidade do material de entrada usado no processo do que estiver disponível do inventário, porém as quantidades das substâncias produzidas pela destilação não correspondem à um percentual do material de entrada como era de se esperar!

Em `targetMaterial.ReduceComponent(ext.substance, ext.quantity)` estamos subtraindo o percentual da quantidade!

Se o material Carvão é composto de 
80% Carbono
5% Enxofre
15% Metano
Uma unidade de Carvão queimada deveria retornar
.8 de Carbono, .05 de Enxofre e .15 de Metano!

A `func (m *Machine) calculateExtraction(d Refine, dt float32)` está misturando os percentuais das substâncias que compõem o material com quantidade no momento de calcular as perdas!

Isso ficou claro quando refatoramos o termo `quantity` em `Composite` para `percentual`!

Todos os pontos onde `Composite.percentual` é usado, devemos reconsiderar como é usado e passar a tratá-lo como o percentual que ele representa. Para facilitar o tratamento de quantidades, devemos considerar que o jogo trata materiais como unidades, portanto temos sempre **uma unidade de carvão** ou **uma unidade de ferro** nos inventários quando são produzidas ou consumidas. E o cálculo de perda, devemos calcular `maxLoss` na `func (m *Machine) calculateExtraction(d Refine, dt float32)` como percentual!

Para isso na fórmula `currentLoss := (deltaTemp * element.Volatility) * dt`, currentLoss deve ser a proporção de material que será perdida durante a queima do material, e não sua quantidade absoluta!

Usando a Fornalha:
```go
		ID:           CoalBurnerID,
		Name:         "Fogueira de Acampamento",
		Heat:         310.0, // Iniciada com um fósforo (acima dos 300°C de ativação)
		CurrentChain: []MachineOp{
			CoalBurnProcess, // A fogueira está configurada para queimar carvão
		},
		Progress: 0,
```

No processo de queima de carvão,
```go
	Operation: Operation{
		ID:               CombustionCoalID,
		Name:             "Combustão de Carvão",
		RequiredTemp:     600.0, // Temperatura ideal de queima
		ActivationTemp:   300.0, // FlashPoint do seu Carvão
		Duration:         10.0,  // Cada unidade de carvão dura 10 segundos
		EnergyCost:       0,     // Não gasta energia elétrica/externa
		BaseEnergyChange: 150.0, // Produz calor para a máquina e arredores
	},
	// O que a fogueira "consome" do mundo
	Input: CoalID,
	// O que sobra (Cinzas/Carbono e Gases para a atmosfera)
	Outputs: []Composite{
		{substance: PureCarbonID, percentual: 5}, // Sobra um pouco de cinza/resíduo
		{substance: MethaneID, percentual: 15},   // Libera os voláteis
		{substance: PureSulfurID, percentual: 5}, // Libera enxofre (poluição)
	},
```

Quando pegamos um material como o Carvão:
```go
		ID:   CoalID,
		Name: "Coal",
		Composites: []Composite{
			{PureCarbonID, 80},
			{PureSulfurID, 5},
			{MethaneID, 15},
		},
		State:         Solid,
		HP:            80,    // Carvão é quebradiço
		Temperature:   21,    // Temp ambiente inicial
		Hardness:      2.5,   // Fácil de minerar
		Density:       1.5,   // Leve
		SpecificHeat:  0.9,   // Esquenta relativamente rápido
		EnergyDensity: 24.0,  // Valor alto para combustão (MJ/kg)
		FlashPoint:    300.0, // Acende com calor moderado
		Conductivity:  0.2,   // Pobre condutor (ajuda a manter o calor na fornalha)
```

Para o Carbono:
```go
    ID:           CarbonID, 
    Nanme:        "Carbon", 
    Symbol:       "C", 
    Weight:       12.011, 
    Category:     "Nonmetal", 
    IsRadioactive: false, 
    FusionFuel:    false, 
    BoilingPoint:  4827, 
    Volatility:    0.05
```

Feito o ajuste nas funções de armazenamento no inventário das máquinas `ProduceSubstance` e `ProduceMaterial` para converter a quantidade informada em unidade fracionária, o sistema começou a ficar coerente, apesar de que a perda ainda não está sendo considerada!

Para isso devemos começar a usar a estrutura `Outputs`, que seriam as saídas dos processos!

Essa estrutura deveria ser considerada em `func (m *Machine) finishDistillation(d Refine, dt float32)` de [refine.go](./pkg/craft/refine.go) que atualmente apenas atualiza os inventários com as quantidades brutas, ou seja, a estrutura `Refine` não está sendo usada corretamente! `Outputs` lista as perdas do processo, e estas perdas devem ser jogadas no inventário!

A idéia é consumir `Input` do inventário e depositar os `Outputs` restantes no mesmo inventário e está faltando só fazer isso!