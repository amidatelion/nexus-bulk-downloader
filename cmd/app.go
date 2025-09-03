package cmd

func Run() {
	RootCmd.AddCommand(DownloadCmd)
	RootCmd.Execute()
}
