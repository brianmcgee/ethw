package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tyler-smith/go-bip39"
)

func TestUnmarshalText_ValidSeedAndAlias(t *testing.T) {
	sd := &SeedData{}
	seed := "myvalidseed"
	alias := "myalias"

	seedBytes := []byte("seed=" + seed + ";alias=" + alias)
	err := sd.UnmarshalText(seedBytes)

	assert.Nil(t, err)
	assert.Equal(t, bip39.NewSeed(seed, ""), sd.Data)
	assert.Equal(t, alias, sd.Alias)
}

func TestUnmarshalText_ValidSeedNoAlias(t *testing.T) {
	sd := &SeedData{}
	seed := "myvalidseed"

	seedBytes := []byte("seed=" + seed)
	err := sd.UnmarshalText(seedBytes)

	assert.Nil(t, err)
	assert.Equal(t, bip39.NewSeed(seed, ""), sd.Data)
	assert.Empty(t, sd.Alias)
}

func TestUnmarshalText_EmptySeed(t *testing.T) {
	sd := &SeedData{}

	seedBytes := []byte("seed=")
	err := sd.UnmarshalText(seedBytes)

	assert.Equal(t, ErrInvalidSeedFormat, err)
}

func TestUnmarshalText_EmptyAlias(t *testing.T) {
	sd := &SeedData{}
	seed := "myvalidseed"

	seedBytes := []byte("seed=" + seed + ";alias=")
	err := sd.UnmarshalText(seedBytes)

	assert.Nil(t, err)
	assert.Equal(t, bip39.NewSeed(seed, ""), sd.Data)
	assert.Empty(t, sd.Alias)
}

func TestUnmarshalText_InvalidFormat(t *testing.T) {
	sd := &SeedData{}

	seedBytes := []byte("invaliddata")
	err := sd.UnmarshalText(seedBytes)

	assert.Equal(t, ErrInvalidSeedFormat, err)
}
