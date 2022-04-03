//go:build !freebsd || (freebsd && cgo)
// +build !freebsd freebsd,cgo

package cmd

import (
	"github.com/99designs/keyring"
	"github.com/spf13/cobra"
)

type secretKeyringCmdConfig struct {
	delete secretKeyringDeleteCmdConfig
	get    secretKeyringGetCmdConfig
	set    secretKeyringSetCmdConfig
}

type secretKeyringDeleteCmdConfig struct {
	service string
	user    string
}

type secretKeyringGetCmdConfig struct {
	service string
	user    string
}

type secretKeyringSetCmdConfig struct {
	service string
	user    string
	value   string
}

func (c *Config) newSecretKeyringCmd() *cobra.Command {
	keyringCmd := &cobra.Command{
		Use:   "keyring",
		Args:  cobra.NoArgs,
		Short: "Interact with keyring",
	}

	keyringDeleteCmd := &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "Delete a value from keyring",
		RunE:  c.runSecretKeyringDeleteCmdE,
	}
	secretKeyringDeletePersistentFlags := keyringDeleteCmd.PersistentFlags()
	secretKeyringDeletePersistentFlags.StringVar(&c.secret.keyring.get.service, "service", "", "service")
	secretKeyringDeletePersistentFlags.StringVar(&c.secret.keyring.get.user, "user", "", "user")
	markPersistentFlagsRequired(keyringDeleteCmd, "service", "user")
	keyringCmd.AddCommand(keyringDeleteCmd)

	keyringGetCmd := &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "Get a value from keyring",
		RunE:  c.runSecretKeyringGetCmdE,
	}
	secretKeyringGetPersistentFlags := keyringGetCmd.PersistentFlags()
	secretKeyringGetPersistentFlags.StringVar(&c.secret.keyring.get.service, "service", "", "service")
	secretKeyringGetPersistentFlags.StringVar(&c.secret.keyring.get.user, "user", "", "user")
	markPersistentFlagsRequired(keyringGetCmd, "service", "user")
	keyringCmd.AddCommand(keyringGetCmd)

	keyringSetCmd := &cobra.Command{
		Use:   "set",
		Args:  cobra.NoArgs,
		Short: "Set a value in keyring",
		RunE:  c.runSecretKeyringSetCmdE,
	}
	secretKeyringSetPersistentFlags := keyringSetCmd.PersistentFlags()
	secretKeyringSetPersistentFlags.StringVar(&c.secret.keyring.set.service, "service", "", "service")
	secretKeyringSetPersistentFlags.StringVar(&c.secret.keyring.set.user, "user", "", "user")
	secretKeyringSetPersistentFlags.StringVar(&c.secret.keyring.set.value, "value", "", "value")
	markPersistentFlagsRequired(keyringSetCmd, "service", "user")
	keyringCmd.AddCommand(keyringSetCmd)

	return keyringCmd
}

func (c *Config) runSecretKeyringDeleteCmdE(cmd *cobra.Command, args []string) error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: c.secret.keyring.delete.service,
	})
	if err != nil {
		return err
	}
	return ring.Remove(c.secret.keyring.delete.user)
}

func (c *Config) runSecretKeyringGetCmdE(cmd *cobra.Command, args []string) error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: c.secret.keyring.get.service,
	})
	if err != nil {
		return err
	}
	item, err := ring.Get(c.secret.keyring.get.user)
	if err != nil {
		return err
	}
	return c.writeOutput(item.Data)
}

func (c *Config) runSecretKeyringSetCmdE(cmd *cobra.Command, args []string) error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: c.secret.keyring.set.service,
	})
	if err != nil {
		return err
	}
	data := c.secret.keyring.set.value
	if data == "" {
		var err error
		data, err = c.readPassword("Value: ")
		if err != nil {
			return err
		}
	}
	return ring.Set(keyring.Item{
		Key:  c.secret.keyring.set.user,
		Data: []byte(data),
	})
}
