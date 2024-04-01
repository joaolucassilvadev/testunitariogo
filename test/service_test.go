package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	//	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"test.com/controller"
)

func TestService(t *testing.T) {
	// Bom o mockresponse é uma variavel a qual estamos guardando a resposta esperada para rota especifica
	mockResponse := `{"mensage":"Welcome to the mato"}`
	// Bom aqui o r esta sendo atribuido a minha funcao a onde eu injeto o Engine assim podendo simular minhas rotas por ele é o main_test.go
	r := SetUpRouter()
	// Aqui eu coloco meu r com get como se foi atribuido ao meu r o gin.Engine logo ele vira minha rota
	r.GET("/", controller.Testservice)
	// Bom esse req é uma simplificação da palavra request mas poderiamos colocar qqualquer nome é mais pela nomeclatura
	// o http.Newrequest é uma funcao da biblioteca net/http, essa funcao serve para fazermos uma nova requisição ou um request a nossa rota
	// ela espera 3 argumentos o tipo de rota que é o get a onde ela vai fazer a solicitação e o ultimo é saber se ela tem corpo logo não tem então botamos nil
	req, _ := http.NewRequest("GET", "/", nil)
	// bom apos fazermos nossa nova request obtemos a saida logo vamos gravar essa saida a nosso w com o httptest.NewRecord, ela faz parte da nossa biblioteca padrão do para test
	w := httptest.NewRecorder()
	// Bom aqui temos o http.ResponseWriter que é usado para tratar uma solicitação http ela recebe dois argumentos
	// um objeto http.ResponseWriter que é usado para escrever a resposta http de volta ao cliente
	// um objeto http.request que representa a solicitação http recebida pelo servidor
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
