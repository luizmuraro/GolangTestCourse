package main

type Dicionario map[string]string

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra que você procurou")
	ErrpalavraExistente   = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrpalavraInexistente = ErrDicionario("não foi possívle atualizar a palavra pois ela não existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(palavra string) (string, error) {

	palavra, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return palavra, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrpalavraExistente
	default:
		return err

	}

	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		return ErrpalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}
	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}
