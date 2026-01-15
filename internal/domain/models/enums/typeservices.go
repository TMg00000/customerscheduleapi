package enums

type TypeService int

const (
	Brasileiro TypeService = iota
	FioAFio
	EfeitoDelineado
	Egipcio
	Fox
	Gatinho
	Tufo
	Russo
)

func (ts TypeService) messages() string {
	switch ts {
	case Brasileiro:
		return "Brasileiro"
	case FioAFio:
		return "Fio a Fio"
	case EfeitoDelineado:
		return "Efeito Delineado"
	case Egipcio:
		return "Egípcio"
	case Fox:
		return "Fox"
	case Gatinho:
		return "Gatinho"
	case Tufo:
		return "Tufo"
	case Russo:
		return "Russo"
	default:
		return ""
	}
}
