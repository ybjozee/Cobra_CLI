package cmd

import (
	"cipher_cli/pkg/bacon"
	"cipher_cli/pkg/caesar"
	"cipher_cli/pkg/notification"
	"cipher_cli/pkg/validation"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a message",
	Long: `Use this command to generate a cryptic version of a message by providing the text you want to encrypt
and the encryption algorithm to be applied. For example:

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=caesar --key=54

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=bacon
`,
	Run: func(cmd *cobra.Command, args []string) {
		plaintext := strings.Join(args, " ")
		var ciphertext = ""
		algorithm := cmd.Flags().Lookup("algorithm").Value.String()
		key := cmd.Flags().Lookup("key").Value.String()
		recipient := cmd.Flags().Lookup("recipient").Value.String()

		if strings.ToLower(algorithm) == CAESAR {
			rotations, err := validation.GetRotationsFromKey(key)
			if err != nil {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}
			ciphertext = caesar.Encrypt(plaintext, rotations)
		} else {
			ciphertext = bacon.Encrypt(plaintext)
		}
		fmt.Printf("Plaintext: %s\nCiphertext: %s\n", plaintext, ciphertext)
		if recipient != "" {
			err := validation.ValidatePhoneNumber(recipient)
			if err == nil {
				notification.SendMessage(recipient, fmt.Sprintf("From your partner in mischief\n%s", ciphertext))
			} else {
				fmt.Printf("Error: %s", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	//TODO: Flag to randomize casing
	encryptCmd.Flags().StringP("algorithm", "a", "", "The algorithm to use for this action")
	_ = encryptCmd.MarkFlagRequired("algorithm")
	encryptCmd.Flags().StringP("recipient", "r", "", "Send encrypted messages to a phone number")
	rootCmd.AddCommand(encryptCmd)
}
