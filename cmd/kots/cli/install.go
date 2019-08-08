package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/replicatedhq/kots/pkg/kotsadm"
	"github.com/replicatedhq/kots/pkg/pull"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "install [upstream uri]",
		Short:         "",
		Long:          ``,
		SilenceUsage:  true,
		SilenceErrors: false,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}

			rootDir, err := ioutil.TempDir("", "kotsadm")
			if err != nil {
				return err
			}
			defer os.RemoveAll(rootDir)

			pullOptions := pull.PullOptions{
				HelmRepoURI: v.GetString("repo"),
				RootDir:     rootDir,
				Overwrite:   false,
				Namespace:   v.GetString("namespace"),
			}
			if err := pull.Pull(args[0], pullOptions); err != nil {
				return err
			}

			deployOptions := kotsadm.DeployOptions{
				Namespace:      v.GetString("namespace"),
				Kubeconfig:     v.GetString("kubeconfig"),
				IncludeShip:    v.GetBool("include-ship"),
				IncludeGitHub:  v.GetBool("include-github"),
				SharedPassword: v.GetString("shared-password"),
			}
			if err := kotsadm.Deploy(deployOptions); err != nil {
				return err
			}

			// deploy kotsadm to the namespace

			// upload the kots app to kotsadm

			return nil
		},
	}

	cmd.Flags().String("kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "the kubeconfig to use")
	cmd.Flags().String("namespace", "default", "the namespace to deploy to")
	cmd.Flags().Bool("include-ship", false, "include the shipinit/edit/update and watch components")
	cmd.Flags().Bool("include-github", false, "set up for github login")
	cmd.Flags().String("shared-password", "", "shared password to apply")

	cmd.Flags().String("repo", "", "repo uri to use when installing a helm chart")
	cmd.Flags().StringArray("set", []string{}, "values to pass to helm when running helm template")

	return cmd
}
