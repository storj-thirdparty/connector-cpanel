package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command.
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command to upload data to storj V3 network.",
	Long:  `Command to connect and transfer back-up file from a desired cPanel instance to given Storj Bucket.`,
	Run:   influxStore,
}

func init() {

	// Setup the store command with its flags.
	rootCmd.AddCommand(storeCmd)
	var defaultInfluxFile string
	var defaultStorjFile string
	storeCmd.Flags().BoolP("accesskey", "a", false, "Connect to storj using access key(default connection method is by using API Key).")
	storeCmd.Flags().BoolP("share", "s", false, "For generating share access of the uploaded backup file.")
	storeCmd.Flags().StringVarP(&defaultInfluxFile, "cpanel", "c", "./config/cpanel_property.json", "full filepath contaning cpanel configuration.")
	storeCmd.Flags().StringVarP(&defaultStorjFile, "storj", "u", "./config/storj_config.json", "full filepath contaning storj V3 configuration.")
}

func influxStore(cmd *cobra.Command, args []string) {

	// Process arguments from the CLI.
	cPanelConfigfilePath, _ := cmd.Flags().GetString("cpanel")
	fullFileNameStorj, _ := cmd.Flags().GetString("storj")
	useAccessKey, _ := cmd.Flags().GetBool("accesskey")
	useAccessShare, _ := cmd.Flags().GetBool("share")

	// Read cpanel instance's configurations from an external file and create an cpanel configuration object.
	configcPanel := LoadcPanelProperty(cPanelConfigfilePath)

	// Read storj network configurations from and external file and create a storj configuration object.
	storjConfig := LoadStorjConfiguration(fullFileNameStorj)

	// Connect to storj network using the specified credentials.
	access, project := ConnectToStorj(storjConfig, useAccessKey)

	// Connect to cPanel instance, create back-up and get information of the back-up file to be uploaded.
	reader := ConnectToCpanel(configcPanel)

	// Uplaod the cpanel backup file to storj network
	fmt.Printf("\nInitiating back-up.\n")
	UploadData(project, storjConfig, reader.FileName, reader.FileHandle)
	fmt.Printf("\nBack-up complete.\n\n")

	// Create restricted shareable serialized access if share is provided as argument.
	if useAccessShare {
		ShareAccess(access, storjConfig)
	}
}
