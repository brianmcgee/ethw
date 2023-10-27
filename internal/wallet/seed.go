package wallet

import (
	"errors"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/tyler-smith/go-bip39"
)

// ErrInvalidSeedFormat represents an error when the seed string is invalid.
var ErrInvalidSeedFormat = errors.New("invalid seed format")

// seedDataRegex is a compiled regular expression for parsing seed strings.
var seedDataRegex = regexp.MustCompile(`(?:seed=([^;]*);?)?(?:alias=([^;]*);?)?`)

// SeedData represents the seed information for a wallet.
type SeedData struct {
	Alias string
	Data  []byte
}

// UnmarshalText unmarshals the SeedData from text
func (sd *SeedData) UnmarshalText(text []byte) error {
	matches := seedDataRegex.FindStringSubmatch(string(text))
	if matches == nil || matches[1] == "" {
		return ErrInvalidSeedFormat
	}

	sd.Data = bip39.NewSeed(strings.TrimSpace(matches[1]), "")
	sd.Alias = strings.TrimSpace(matches[2])

	log.Debugf("Decoded SeedData: %+v", sd)

	return nil
}
