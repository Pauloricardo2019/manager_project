package token

import tokenRepo "gerenciador/adapter/database/sql/token"

func ValidateToken(token string, getTokenByValue tokenRepo.GetTokenByValueFn) (bool, uint64) {

	found, tokenFind, err := getTokenByValue(token)
	if err != nil {
		return false, 0
	}
	if !found {
		return false, 0
	}

	if tokenFind.CompanyID <= 0 {
		return false, 0
	}

	return true, tokenFind.CompanyID
}
