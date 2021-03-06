// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package talos

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/talos-systems/crypto/x509"

	"github.com/talos-systems/talos/internal/pkg/tui/installer"
	machineapi "github.com/talos-systems/talos/pkg/machinery/api/machine"
	"github.com/talos-systems/talos/pkg/machinery/client"
)

var applyConfigCmdFlags struct {
	certFingerprints []string
	filename         string
	insecure         bool
	interactive      bool
	noReboot         bool
}

// applyConfigCmd represents the applyConfiguration command.
var applyConfigCmd = &cobra.Command{
	Use:   "apply-config",
	Short: "Apply a new configuration to a node",
	Long:  ``,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			cfgBytes []byte
			e        error
		)

		if applyConfigCmdFlags.filename != "" {
			cfgBytes, e = ioutil.ReadFile(applyConfigCmdFlags.filename)
			if e != nil {
				return fmt.Errorf("failed to read configuration from %q: %w", applyConfigCmdFlags.filename, e)
			}

			if len(cfgBytes) < 1 {
				return fmt.Errorf("no configuration data read")
			}
		} else if !applyConfigCmdFlags.interactive {
			return fmt.Errorf("no filename supplied for configuration")
		}

		withClient := func(f func(context.Context, *client.Client) error) error {
			if applyConfigCmdFlags.insecure {
				ctx := context.Background()

				if len(Nodes) != 1 {
					return fmt.Errorf("insecure mode requires one and only one node, got %d", len(Nodes))
				}

				c, err := client.New(ctx, client.WithTLSConfig(&tls.Config{
					InsecureSkipVerify: true,
				}), client.WithEndpoints(Nodes...))
				if err != nil {
					return err
				}

				//nolint: errcheck
				defer c.Close()

				tlsConfig := &tls.Config{
					InsecureSkipVerify: true,
				}

				if len(applyConfigCmdFlags.certFingerprints) > 0 {
					fingerprints := make([]x509.Fingerprint, len(applyConfigCmdFlags.certFingerprints))

					for i, stringFingerprint := range applyConfigCmdFlags.certFingerprints {
						fingerprints[i], err = x509.ParseFingerprint(stringFingerprint)
						if err != nil {
							return fmt.Errorf("error parsing certificate fingerprint %q: %v", stringFingerprint, err)
						}
					}

					tlsConfig.VerifyConnection = x509.MatchSPKIFingerprints(fingerprints...)
				}

				c, err = client.New(ctx, client.WithTLSConfig(tlsConfig), client.WithEndpoints(Nodes...))
				if err != nil {
					return err
				}

				return f(ctx, c)
			}

			return WithClient(f)
		}

		return withClient(func(ctx context.Context, c *client.Client) error {
			if applyConfigCmdFlags.interactive {
				install := installer.NewInstaller()
				node := Nodes[0]

				return WithClientNoNodes(func(bootstrapCtx context.Context, bootstrapClient *client.Client) error {
					opts := []installer.Option{}

					if len(Endpoints) > 0 {
						opts = append(opts, installer.WithBootstrapNode(bootstrapCtx, bootstrapClient, Endpoints[0]))
					}

					conn, err := installer.NewConnection(
						ctx,
						c,
						node,
						opts...,
					)
					if err != nil {
						return err
					}

					return install.Run(conn)
				})
			}

			if _, err := c.ApplyConfiguration(ctx, &machineapi.ApplyConfigurationRequest{
				Data:     cfgBytes,
				NoReboot: applyConfigCmdFlags.noReboot,
			}); err != nil {
				return fmt.Errorf("error applying new configuration: %s", err)
			}

			return nil
		})
	},
}

func init() {
	applyConfigCmd.Flags().StringVarP(&applyConfigCmdFlags.filename, "file", "f", "", "the filename of the updated configuration")
	applyConfigCmd.Flags().BoolVarP(&applyConfigCmdFlags.insecure, "insecure", "i", false, "apply the config using the insecure (encrypted with no auth) maintenance service")
	applyConfigCmd.Flags().StringSliceVar(&applyConfigCmdFlags.certFingerprints, "cert-fingerprint", nil, "list of server certificate fingeprints to accept (defaults to no check)")
	applyConfigCmd.Flags().BoolVar(&applyConfigCmdFlags.interactive, "interactive", false, "apply the config using text based interactive mode")
	applyConfigCmd.Flags().BoolVar(&applyConfigCmdFlags.noReboot, "no-reboot", false, "apply the config only after the reboot")

	addCommand(applyConfigCmd)
}
