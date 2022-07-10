package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

//公共常量，令牌过期
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is Invalid")
)

//token有效负载数据
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issued_at"`  //建立时间
	ExpiredAt time.Time `json:"expired_at"` //过期时间
}

//创建新的令牌负载，具有特定的用户名和持续时间
func NewPayload(username string, duration time.Duration) (*Payload, error) {

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil

}

//验证token是否过期
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
