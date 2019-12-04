package response

import "time"

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType	 string `json:"token_type"`
	ExpiresIn	 int	`json:"expires_in"`
}

func CreateAuthResponse(accessToken string, refreshToken string, expiredDate time.Time) AuthResponse {
	tokenType := "bearer"
	expiredInt := expiredDate.Sub(time.Now())
	return AuthResponse {
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType: tokenType,
		ExpiresIn: int(expiredInt)}
}