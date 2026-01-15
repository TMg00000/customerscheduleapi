package enums

type TypeService int

const (
	brasileito TypeService = iota
	fioAFio
	efeitoDelineado
	egipcio
	fox
	gatinho
	tufo
	russo
)

func (ts TypeService) messages() string {
	switch ts {
	case brasileito:
		return "Brasileiro"
	case fioAFio:
		return "Fio a Fio"
	case efeitoDelineado:
		return "Efeito Delineado"
	case egipcio:
		return "Egípcio"
	case fox:
		return "Fox"
	case gatinho:
		return "Gatinho"
	case tufo:
		return "Tufo"
	case russo:
		return "Russo"
	default:
		return ""
	}
}
