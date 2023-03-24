package main

// Refs: https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/ola-mundo

import "testing"

// Para criar um teste as seguintes regras devem ser seguidas
// Criar um arquivo com mesmo nome com final _test.go (ex: ola_test.go)
// A função de test deve começar com Test (ex: Testola)
// A função deve receber um unico argumento, que é t *testing.T
func TestOla(t *testing.T) {
	// A função é atribuida á uma variavel ?
	// t.Helper marca a funçãp chamada como uma função de auxilio
	// dessa forma a mensagem de erro será exibida para o ponto onde o erro ocorreu
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	// o metodo Run executa um subtest.
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		name		:= "Davy"
		resultado	:= Ola(name, "")
		esperado	:= "Olá, " + name
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado	:= Ola("", "")
		esperado	:= "Olá, mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'olá, Greg' em espanhol", func(t *testing.T) {
		name		:= "Greg"
		lang		:= "espanhol"
		resultado	:= Ola(name, lang)
		esperado	:= "Hola, Greg"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'olá, Greg' em francês", func(t *testing.T) {
		name		:= "Greg"
		lang		:= "francês"
		resultado	:= Ola(name, lang)
		esperado	:= "Bonjour, Greg"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'olá, Greg' em japonês", func(t *testing.T) {
		name		:= "Gureggu"
		lang		:= "japonês"
		resultado	:= Ola(name, lang)
		esperado	:= "Kon'nichiwa, Gureggu"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}