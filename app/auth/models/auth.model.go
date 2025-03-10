package models

type User struct {
    ID       uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
    Username string `json:"username,omitempty"`
    Password string `json:"password,omitempty"`
	Hp string `json:"hp,omitempty"`
	Address string `json:"address,omitempty"`
	Bank Bank `json:"bank"`
	Pocket Pocket `json:"pocket"`
	Term Term `json:"term"`
	TotalSaldo int `json:"total_saldo,omitempty"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Name string `json:"name"`
    Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Hp string `json:"hp"`
	Address string `json:"address"`
    Token string `json:"token"`
}

type Bank struct {
    Name string `json:"name"`
    Account int `json:"account"`
}

type Pocket struct {
    Saldo int `json:"saldo"`
}

type Term struct {
    PrincipalDeposit int `json:"principal_deposit"`
    DepositInterestProfit int `json:"deposit_interest_profit"`
	DepositInterestTax int `json:"deposit_interest_tax"`
    TotalInvestment int `json:"total_investment"`
}
