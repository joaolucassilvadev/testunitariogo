package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"test.com/model"
)

func Popoluar() {
	log.Printf("inicio da funcao popular")
	var Servidor model.Servidor
	file, err := os.Open("banco_servidor_funben.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&model.Servidor{})

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Inserção de dados concluída com sucesso.")
				break
			}
			fmt.Println("Erro ao ler o CSV:", err)
			return
		}

		Servidor.Canal_solicitacao_beneficio = record[0]
		Servidor.Tipo_inclusao = record[1]
		Servidor.Secretaria_orgao = "Emserh"
		Servidor.N_do_processo = record[3]
		Servidor.Data_abertura_do_processo = record[4]
		Servidor.Cpf_do_titular, _ = strconv.Atoi(record[5])
		Servidor.Matricula_titular, _ = strconv.Atoi(record[6])
		Servidor.Nome_do_titular = record[7]
		Servidor.Data_nascimento_do_titular = record[8]
		Servidor.Sexo = record[9]
		Servidor.Data_inclusao_na_folha_funben = record[10]
		Servidor.Lotacao_setor = record[11]
		Servidor.Data_de_admissao_emserh = record[12]

		db.Create(&Servidor)
	} //
}
