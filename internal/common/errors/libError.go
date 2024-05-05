package customError

import "errors"

var (
	ErrTokenWithClaimsIsNil = errors.New("token with claims is nil")
	ErrGenerateAccessToken  = errors.New("failed to generate access token")
	ErrGenerateRefreshToken = errors.New("failed to generate refresh token")
)
