package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserInfo struct {
	Cpf          string  `json:"cpf"`
	BancoDaConta string  `json:"banco_da_conta"`
	QuantiaSaldo float64 `json:"quantia_saldo"`
	Emprestimos  []struct {
		InstituicaoFinanceira string  `json:"instituicao_financeira"`
		Produto               string  `json:"produto"`
		Quantia               float64 `json:"quantia"`
		DataFirmacao          string  `json:"data_firmacao"`
		DataEncerramento      string  `json:"data_encerramento"`
		Periodicidade         string  `json:"periodicidade"`
		TaxaDeJuros           float64 `json:"taxa_de_juros"`
		TipoDeGarantia        string  `json:"tipo_de_garantia"`
		QuantiaSobGarantia    float64 `json:"quantia_sob_garantia"`
	} `json:"emprestimos"`
	Transacoes []struct {
		Banco            string  `json:"banco"`
		TipoTransacao    string  `json:"tipo_transacao"`
		QuantiaTransacao float64 `json:"quantia_transacao"`
		DataTransacao    string  `json:"data_transacao"`
	} `json:"transacoes"`
	Cartoes []struct {
		NomeCartao        string  `json:"nome_cartao"`
		BancoCartao       string  `json:"banco_cartao"`
		EmissoraCartao    string  `json:"emissora_cartao"`
		Limite            float64 `json:"limite"`
		LimiteUtilizado   float64 `json:"limite_utilizado"`
		FaturasAnteriores []struct {
			DataVencimentoFatura string  `json:"data_vencimento_fatura"`
			ValorFatura          float64 `json:"valor_fatura"`
		} `json:"faturas_anteriores"`
		Transacoes []struct {
			DataTransacao  string  `json:"data_transacao"`
			ValorTransacao float64 `json:"valor_transacao"`
			TipoTransacao  string  `json:"tipo_transacao"`
		} `json:"transacoes"`
	} `json:"cartoes"`
}
