package pet_airbnb

import "errors"

type (
	Sitter struct {
		PhoneNumber string
		Email       string
		Verified    int
	}
	verifier struct{}
	Verifier interface {
		Start(*Sitter) error
		Complete(*Sitter) error
	}
	Customer struct{}
	Review   struct{}
	Photo    struct{}
	Profile  struct{}
)

const (
	NOT_VERIFIED int = iota
	IN_PROGRESS
	VERIFIED
)

var (
	ErrInvalidState = errors.New("invalid state")
)

type verifierMock struct {
}

func (v *verifierMock) Start(s *Sitter) error {
	if s.Verified == NOT_VERIFIED {
		s.Verified = IN_PROGRESS
		return nil
	}
	return ErrInvalidState
}

func (v *verifierMock) Complete(s *Sitter) error {
	if s.Verified == IN_PROGRESS {
		s.Verified = VERIFIED
		return nil
	}
	return ErrInvalidState
}
