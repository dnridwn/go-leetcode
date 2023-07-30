package design_authentication_manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TOKEN_ID_TEST string = "ThisIsToken"

func TestAuthenticationManager_Generate(t *testing.T) {
	am := Constructor(5)
	am.Generate(TOKEN_ID_TEST, 0)

	assert.Equal(t, &Token{
		TokenId:   TOKEN_ID_TEST,
		CreatedAt: 0,
		ExpiredAt: 5,
	}, am.token[TOKEN_ID_TEST])
}

func TestAuthenticationManager_Renew(t *testing.T) {
	testCases := []struct {
		name              string
		createAuthManager func() AuthenticationManager
		tokenId           string
		currentTime       int
		expected          *Token
	}{
		{
			name: "Test Renew Unexpired Token",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token: map[string]*Token{
						TOKEN_ID_TEST: {
							TokenId:   TOKEN_ID_TEST,
							CreatedAt: 0,
							ExpiredAt: 3,
						},
					},
				}
			},
			tokenId:     TOKEN_ID_TEST,
			currentTime: 2,
			expected: &Token{
				TokenId:   TOKEN_ID_TEST,
				CreatedAt: 0,
				ExpiredAt: 5,
			},
		},
		{
			name: "Test Renew Expired Token",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token: map[string]*Token{
						TOKEN_ID_TEST: {
							TokenId:   TOKEN_ID_TEST,
							CreatedAt: 0,
							ExpiredAt: 3,
						},
					},
				}
			},
			tokenId:     TOKEN_ID_TEST,
			currentTime: 3,
			expected: &Token{
				TokenId:   TOKEN_ID_TEST,
				CreatedAt: 0,
				ExpiredAt: 3,
			},
		},
		{
			name: "Test Renew Non Exist Token",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token:      map[string]*Token{},
				}
			},
			tokenId:     TOKEN_ID_TEST,
			currentTime: 2,
			expected:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			am := tc.createAuthManager()
			am.Renew(tc.tokenId, tc.currentTime)
			assert.Equal(t, tc.expected, am.token[tc.tokenId])
		})
	}
}

func TestAuthenticationManager_CountUnexpiredTokens(t *testing.T) {
	testCases := []struct {
		name              string
		createAuthManager func() AuthenticationManager
		currentTime       int
		expected          int
	}{
		{
			name: "Test All Tokens Are Unexpired",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 4,
					token: map[string]*Token{
						TOKEN_ID_TEST + "_1": {
							TokenId:   TOKEN_ID_TEST + "_1",
							CreatedAt: 0,
							ExpiredAt: 4,
						},
						TOKEN_ID_TEST + "_2": {
							TokenId:   TOKEN_ID_TEST + "_2",
							CreatedAt: 0,
							ExpiredAt: 4,
						},
						TOKEN_ID_TEST + "_3": {
							TokenId:   TOKEN_ID_TEST + "_3",
							CreatedAt: 0,
							ExpiredAt: 4,
						},
					},
				}
			},
			currentTime: 3,
			expected:    3,
		},
		{
			name: "Test Some Tokens Are Unexpired",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token: map[string]*Token{
						TOKEN_ID_TEST + "_1": {
							TokenId:   TOKEN_ID_TEST + "_1",
							CreatedAt: 0,
							ExpiredAt: 5,
						},
						TOKEN_ID_TEST + "_2": {
							TokenId:   TOKEN_ID_TEST + "_2",
							CreatedAt: 0,
							ExpiredAt: 7,
						},
						TOKEN_ID_TEST + "_3": {
							TokenId:   TOKEN_ID_TEST + "_3",
							CreatedAt: 0,
							ExpiredAt: 8,
						},
					},
				}
			},
			currentTime: 6,
			expected:    2,
		},
		{
			name: "Test All Tokens Are Expired",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token: map[string]*Token{
						TOKEN_ID_TEST + "_1": {
							TokenId:   TOKEN_ID_TEST + "_1",
							CreatedAt: 0,
							ExpiredAt: 3,
						},
						TOKEN_ID_TEST + "_2": {
							TokenId:   TOKEN_ID_TEST + "_2",
							CreatedAt: 0,
							ExpiredAt: 3,
						},
						TOKEN_ID_TEST + "_3": {
							TokenId:   TOKEN_ID_TEST + "_3",
							CreatedAt: 0,
							ExpiredAt: 3,
						},
					},
				}
			},
			currentTime: 6,
			expected:    0,
		},
		{
			name: "Test With No Tokens",
			createAuthManager: func() AuthenticationManager {
				return AuthenticationManager{
					timeToLive: 3,
					token:      map[string]*Token{},
				}
			},
			currentTime: 6,
			expected:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			am := tc.createAuthManager()
			assert.Equal(t, tc.expected, am.CountUnexpiredTokens(tc.currentTime))
		})
	}
}
