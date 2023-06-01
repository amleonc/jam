package jam

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/option"
)

type Option = option.Interface

type JamOption interface {
	Option
	constraint()
}

type optionConstraint struct{ Option }

func (*optionConstraint) constraint() {}

type identWithAlgorithm struct{}
type identWithPrivateKey struct{}
type identWithPublicKey struct{}
type identWithKeySet struct{}
type identWithExtractors struct{}

func WithAlgorithm(a SignatureAlgorithm) JamOption {
	return &optionConstraint{option.New(identWithAlgorithm{}, a)}
}

func WithPrivateKey(k jwk.Key) JamOption {
	return &optionConstraint{option.New(identWithPrivateKey{}, k)}
}

func WithPublicKey(k jwk.Key) JamOption {
	return &optionConstraint{option.New(identWithPublicKey{}, k)}
}

func WithKeySet(ks jwk.Set) JamOption {
	return &optionConstraint{option.New(identWithKeySet{}, ks)}
}

func WithExtractors(extractors ...Extractor) JamOption {
	return &optionConstraint{option.New(identWithExtractors{}, extractors)}
}
