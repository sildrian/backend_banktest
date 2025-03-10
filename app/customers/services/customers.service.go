package services

import (
    "banktest/app/customers/models"
	auth_model "banktest/app/auth/models"
	"banktest/config"
	"fmt"
)

func GetAllUserData() (bool, string, models.CustomerResponse) {
    // enable this if you want use unit test
    // config.InitDB()
	// defer config.DB.Close()

	var userRes models.CustomerResponse
	// Query users from database
    rows, err := config.DB.Query(
        "SELECT c.id, c.name, c.username, c.hp, c.address, b.bank_name, b.bank_account, p.saldo, t.principal_deposit, t.deposit_interest_profit, t.deposit_interest_tax, t.total_investment FROM customers c join bank b on c.id=b.id_cust join pocket p on c.id=p.id_cust join term t on c.id=t.id_cust",
    )
    if err != nil {
		fmt.Println(err)
        return false,"Database Error1",userRes
    }
    defer rows.Close()

	// Prepare slice for users
    var users []auth_model.User
	var bank auth_model.Bank
	var pocket auth_model.Pocket
	var term auth_model.Term

    // Iterate through rows
    for rows.Next() {
		var user auth_model.User
        err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Hp, &user.Address, &bank.Name, &bank.Account, &pocket.Saldo, &term.PrincipalDeposit, &term.DepositInterestProfit, &term.DepositInterestTax, &term.TotalInvestment)
        if err != nil {
            return false,"Database Error2",userRes
        }
		user.Bank = bank
		user.Pocket = pocket
		user.Term = term
		user.TotalSaldo = pocket.Saldo + term.TotalInvestment
        users = append(users, user)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return false, "Database Error3", userRes
    }

    // If no users found
    if len(users) == 0 {
        return true, "No Data Found", userRes
    }

    userRes.User = users
    return true,"Success Get User",userRes
}

func GetUserData(user_req models.CustomerRequest) (bool, string, models.CustomerResponse) {
	var userRes models.CustomerResponse
	// Query users from database with LIKE
    rows, err := config.DB.Query(
		"SELECT c.id, c.name, c.username, c.hp, c.address, b.bank_name, b.bank_account, p.saldo, t.principal_deposit, t.deposit_interest_profit, t.deposit_interest_tax, t.total_investment FROM customers c join bank b on c.id=b.id_cust join pocket p on c.id=p.id_cust join term t on c.id=t.id_cust WHERE c.username ILIKE $1",
        "%" + user_req.Username + "%",
    )
    if err != nil {
        return false,"Database Error1",userRes
    }
    defer rows.Close()

	// Prepare slice for users
    var users []auth_model.User
	var bank auth_model.Bank
	var pocket auth_model.Pocket
	var term auth_model.Term

    // Iterate through rows
    for rows.Next() {
		var user auth_model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Hp, &user.Address, &bank.Name, &bank.Account, &pocket.Saldo, &term.PrincipalDeposit, &term.DepositInterestProfit, &term.DepositInterestTax, &term.TotalInvestment)
        if err != nil {
            return false,"Database Error2",userRes
        }
		user.Bank = bank
		user.Pocket = pocket
		user.Term = term
		user.TotalSaldo = pocket.Saldo + term.TotalInvestment
        users = append(users, user)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        return false, "Database Error3", userRes
    }

    // If no users found
    if len(users) == 0 {
        return true, "No Data Found", userRes
    }

    userRes.User = users
	// 	if err == sql.ErrNoRows {
	// 		return true,"No Data Found",users
	// 	}
    return true,"User found",userRes
}
