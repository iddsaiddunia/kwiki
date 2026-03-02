package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwiki/kwiki/installer"
	"github.com/kwiki/kwiki/tui"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Interactively select and install tools",
	Run: func(cmd *cobra.Command, args []string) {
		m := tui.New()
		p := tea.NewProgram(m, tea.WithAltScreen())
		result, err := p.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		final := result.(tui.Model)
		selections := final.Selections()
		if len(selections) == 0 {
			fmt.Println("No tools selected. Exiting.")
			return
		}
		installer.Install(selections)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export current tool selections to a file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := "kwiki-env.yaml"
		if len(args) > 0 {
			file = args[0]
		}
		if err := installer.Export(file); err != nil {
			fmt.Printf("Export failed: %v\n", err)
			return
		}
		fmt.Printf("✅ Exported to %s\n", file)
	},
}

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Install tools from an exported file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := "kwiki-env.yaml"
		if len(args) > 0 {
			file = args[0]
		}
		selections, err := installer.Import(file)
		if err != nil {
			fmt.Printf("Import failed: %v\n", err)
			return
		}
		installer.Install(selections)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tools",
	Run: func(cmd *cobra.Command, args []string) {
		installer.ListTools()
	},
}
