package key

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
	"github.com/tyler-smith/go-bip39"
	"github.com/zarbchain/zarb-go/cmd"
	"github.com/zarbchain/zarb-go/keystore/key"
)

// Generate creates a new keyfile and stores the keyfile in the disk
func Generate() func(c *cli.Cmd) {
	return func(c *cli.Cmd) {
		c.Before = func() { fmt.Println(cmd.ZARB) }
		c.Action = func() {

			passphrase := cmd.PromptPassphrase("Passphrase: ", true)
			label := cmd.PromptInput("Label: ")

			entropy, _ := bip39.NewEntropy(128)
			mnemonic, _ := bip39.NewMnemonic(entropy)
			seed := bip39.NewSeed(mnemonic, passphrase)
			keyObj, err := key.KeyFromSeed(seed)
			if err != nil {
				cmd.PrintErrorMsg("Failed to create key from the seed: %v", err)
				return
			}

			keyfilepath := cmd.ZarbKeystoreDir() + keyObj.Address().String() + ".json"
			err = key.EncryptKeyToFile(keyObj, keyfilepath, passphrase, label)
			if err != nil {
				cmd.PrintErrorMsg("Failed to encrypt: %v", err)
				return
			}

			fmt.Println()
			cmd.PrintInfoMsg("Key path: %v", keyfilepath)
			cmd.PrintInfoMsg("Address: %v", keyObj.Address())
			cmd.PrintSuccessMsg("mnemonic: \"%v\"", mnemonic)
		}
	}
}
