package resourceserrorsmessages

const (
	BadRequest                       = "JSON inserido na requisição é inválido"
	NotFound                         = "Cliente nao encontrado!"
	IdInvalid                        = "O id Fornecido não foi Localizado"
	ErrorIdGenerate                  = "Error ao gerar o Id"
	ErrorAddInDataBase               = "Erro ao Salvar no Banco de dados"
	ErroQueryDataBase                = "Erro ao consultar o banco"
	ErrorUpdateInDataBase            = "Erro ao atualizar agendamento"
	ErrorDeleteAppointment           = "Erro ao cancelar Agendamento"
	ErrorNotFoundAppointments        = "Id nao encontrado no banco de dados"
	NameIsNil                        = "O nome não pode ser nulo ou vazio."
	TheNameMustContainFiftyChar      = "O nome deve conter no maximo 50 caracteres"
	TheNameStartedOrFinishWithSpace  = "O nome não pode começar ou terminar com espaços vazios"
	TheNumberMustContainElevenDigits = "O numero deve conter 11 Digitos"
	OptionInvalid                    = "A opção digitada nao é valida"
	DateTimeIsInvalid                = "A data escolhida esta invalida"
	ChooseADateLaterThanToday        = "Escolha uma data superior a data de hoje"
)
