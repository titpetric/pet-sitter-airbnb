package pet_airbnb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerificationStart(t *testing.T) {
	t.Parallel()

	sitter := &Sitter{}
	verifier := verifierMock{}
	err := verifier.Start(sitter)
	assert.NoError(t, err)
	assert.Equal(t, IN_PROGRESS, sitter.Verified)
}

func TestVerificationComplete(t *testing.T) {
	t.Parallel()

	sitter := &Sitter{
		Verified: IN_PROGRESS,
	}
	verifier := verifierMock{}
	err := verifier.Complete(sitter)
	assert.NoError(t, err)
	assert.Equal(t, VERIFIED, sitter.Verified)
}

func TestVerificationInitialState(t *testing.T) {
	t.Parallel()

	sitter := &Sitter{}
	assert.Equal(t, sitter.Verified, NOT_VERIFIED)
}
