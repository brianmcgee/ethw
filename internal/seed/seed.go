package seed

import (
	"errors"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/tyler-smith/go-bip39"
)

// ErrInvalidSeedFormat represents an error when the seed string is invalid.
var ErrInvalidSeedFormat = errors.New("invalid seed format")

// ErrInvalidSeedMnemonic represents an error when the mnemonic list is invalid.
var ErrInvalidSeedMnemonic = errors.New("invalid mnemonic format")

// seedDataRegex is a compiled regular expression for parsing seed strings.
var seedDataRegex = regexp.MustCompile(`(?:seed=([^;]*);?)(?:alias=([^;]*);?)?`)

// SeedData represents the seed information for a wallet.
type SeedData struct {
	Alias string
	Seed  []byte
}

// UnmarshalText unmarshals the SeedData from text
func (sd *SeedData) UnmarshalText(text []byte) error {
	matches := seedDataRegex.FindStringSubmatch(string(text))

	// Check for invalid format or missing seed
	if matches == nil || matches[1] == "" {
		log.Debugf("No match found for the current seed")
		return ErrInvalidSeedFormat
	}

	rawSeed := strings.TrimSpace(matches[1])
	log.Debugf("Raw seed: %s", rawSeed)

	if !bip39.IsMnemonicValid(rawSeed) {
		log.Debugf("Invalid nmnemonic passed")
		return ErrInvalidSeedMnemonic
	}

	sd.Seed = bip39.NewSeed(rawSeed, "")
	sd.Alias = strings.TrimSpace(matches[2])

	log.Debugf("Decoded SeedData: %+v", sd)

	return nil
}
