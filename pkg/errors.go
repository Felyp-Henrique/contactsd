package pkg

import "errors"

var (
	ErrorDatabaseConnectionUnavailable = errors.New("Conexão com banco de dados indisponível!")
	ErrorContactsEmpty                 = errors.New("Não foi possível carregar os contatos.")
	ErrorContactsNotFound              = errors.New("Contato não pode ser encontrado!")
	ErrorContactsInsert                = errors.New("Não foi possível cadastrar o contato! Tente novamente.")
	ErrorContactsUpdate                = errors.New("Não foi possível atualizar o contato! Tente novamente.")
)
