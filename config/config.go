package config

import (
	"os"
	"strconv"
)

var Config = struct {
  CreditAmount int
}{
  CreditAmount: 100, // default amount
}


func init() {
	if amount := os.Getenv("CREDIT_AMOUNT"); amount != "" {
	  Config.CreditAmount, _ = strconv.Atoi(amount) 
	}
  }