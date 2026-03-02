package cmd

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var bannerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
var subStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

var rootCmd = &cobra.Command{
	Use:   "kwiki",
	Short: "kwiki - your cross-platform dev environment setup tool",
	Long:  banner(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner())
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(installCmd, listCmd, exportCmd, importCmd)
}

func banner() string {
	art := `
  ██╗  ██╗██╗    ██╗██╗██╗  ██╗██╗
  ██║ ██╔╝██║    ██║██║██║ ██╔╝██║
  █████╔╝ ██║ █╗ ██║██║█████╔╝ ██║
  ██╔═██╗ ██║███╗██║██║██╔═██╗ ██║
  ██║  ██╗╚███╔███╔╝██║██║  ██╗██║
  ╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝╚═╝  ╚═╝╚═╝`
	sub := fmt.Sprintf("  Your cross-platform dev environment setup tool  |  OS: %s", runtime.GOOS)
	return bannerStyle.Render(art) + "\n" + subStyle.Render(sub) + "\n"
}
