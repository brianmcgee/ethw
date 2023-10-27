package keystore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type KeystoreTestSuite struct {
	suite.Suite
	kst     *KeystoreWrapper
	tempDir string
}

func (suite *KeystoreTestSuite) SetupTest() {
	// Create a temporary directory
	var err error
	suite.tempDir, err = os.MkdirTemp("", "keystore_test")
	if err != nil {
		suite.Fail("Failed to create temp directory", err)
	}

	// Create an instance of keystore with the temporary directory
	suite.kst = NewKeyStore(suite.tempDir)
}

func (suite *KeystoreTestSuite) TearDownTest() {
	// Remove the temporary directory
	if err := os.RemoveAll(suite.tempDir); err != nil {
		suite.Fail("Failed to remove temp directory", err)
	}
}

func (suite *KeystoreTestSuite) TestImportPrivateKey() {
	// Test data
	privateKeyHex := "8e46b439b30731a639a3d94a9016b040a87b3027da8c932af7e1560862d11b58"
	password := "1234"

	// Try importing a private key
	err := suite.kst.ImportPrivateKey(privateKeyHex, password, false)
	assert.NoError(suite.T(), err, "Importing private key should succeed")

	// Validate the account is created
	accounts := suite.kst.Accounts()
	assert.Equal(suite.T(), 1, len(accounts), "One account should exist")

	// Try importing the same key again, should fail as overwrite is false
	err = suite.kst.ImportPrivateKey(privateKeyHex, password, false)
	assert.Error(suite.T(), err, "Importing same private key without overwrite should fail")

	// TODO: Decide if it's worthwhile to implement this method
	// Try importing the same key with overwrite enabled
	// err = suite.kst.ImportPrivateKey(privateKeyHex, password, true)
	// assert.NoError(suite.T(), err, "Importing same private key with overwrite should succeed")
}

// Execute the test suite
func TestKeystoreTestSuite(t *testing.T) {
	suite.Run(t, new(KeystoreTestSuite))
}
