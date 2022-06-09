package pkg

type DefaultError struct {
	Message string `json:"message"`
}

func (e DefaultError) Error() string {
	return e.Message
}

func NewDefaultError(message string) DefaultError {
	return DefaultError{
		Message: message,
	}
}

var (
	ErrorDatabaseConnectionUnavailable = NewDefaultError("Conexão com banco de dados indisponível!")
	ErrorContactsEmpty                 = NewDefaultError("Não foi possível carregar os contatos.")
	ErrorContactsNotFound              = NewDefaultError("Contato não pode ser encontrado!")
	ErrorContactsInsert                = NewDefaultError("Não foi possível cadastrar o contato! Tente novamente.")
	ErrorContactsUpdate                = NewDefaultError("Não foi possível atualizar o contato! Tente novamente.")
	ErrorRequestPostJson               = NewDefaultError("Não foi possível ler os dados enviados.")
)
