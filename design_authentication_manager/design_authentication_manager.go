// https://leetcode.com/problems/design-authentication-manager/

package design_authentication_manager

type Token struct {
	TokenId   string
	CreatedAt int
	ExpiredAt int
}

type AuthenticationManager struct {
	timeToLive int
	token      map[string]*Token
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		timeToLive: timeToLive,
		token:      make(map[string]*Token),
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	am.token[tokenId] = &Token{
		TokenId:   tokenId,
		CreatedAt: currentTime,
		ExpiredAt: currentTime + am.timeToLive,
	}
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, ok := am.token[tokenId]; !ok {
		return
	}
	if am.token[tokenId].ExpiredAt <= currentTime {
		return
	}
	am.token[tokenId].ExpiredAt = currentTime + am.timeToLive
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	counter := 0
	for _, token := range am.token {
		if token.ExpiredAt > currentTime {
			counter += 1
		}
	}
	return counter
}
