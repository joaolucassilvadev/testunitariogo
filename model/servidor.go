package model

type Servidor struct {
	Id                            uint   `json:"Id servidor"`
	Canal_solicitacao_beneficio   string `json:"Canal solicitacao beneficio"`
	Tipo_inclusao                 string `json:"Tipo inclusao"`
	Secretaria_orgao              string `json:"Secretaria orgao"`
	N_do_processo                 string `json:"Numero do processo"`
	Data_abertura_do_processo     string
	Cpf_do_titular                int    `json:"Cpf do titular"`
	Matricula_titular             int    `json:"Matricula Titular"`
	Nome_do_titular               string `json:"Nome do titular"`
	Data_nascimento_do_titular    string `json:" Data de nascimento do titular"`
	Sexo                          string `json:"Sexo"`
	Data_inclusao_na_folha_funben string `json:"Data inclusão na folha funben"`
	Lotacao_setor                 string `json:"Lotacao setor"`
	Data_de_admissao_emserh       string `json:"data de admissao emserh"`
	Status                        string `json:"status"`
}

type Dependente struct {
	Id                            uint   `json:"Id servidor"`
	Canal_solicitacao_beneficio   string `json:"Canal solicitacao beneficio"`
	Tipo_inclusao                 string `json:"Tipo inclusao"`
	Secretaria_orgao              string `json:"Secretaria orgao"`
	N_do_processo                 string `json:"Numero do processo"`
	Data_abertura_do_processo     string
	Cpf_dependente                int    `json:"Cpf do dependente"`
	Nome_do_dependente            string `json:"Nome do dependente"`
	Cpf_do_titular                int    `json:"Cpf do titular"`
	Nome_do_titular               string `json:"Nome do titular"`
	Data_de_nascimento_dependente string `json:"Data de nascimento dependente"`
	Sexo_do_dependente            string `json:"Sexo do dependente"`
	Parentesco                    string `json:"Parentesco"`
	Data_de_inclusao_no_pepleo    string `json:"Data de inclusao peple"`
}
