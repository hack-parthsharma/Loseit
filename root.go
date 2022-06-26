package cmd

import (
	"fmt"
	"os"

	"strconv"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "loseit",
	Short: "Calculate an estimate of daily calorie consumption to meet a weekly weight loss goal",
	Long: `loseit is a simple command line tool that takes a tdee calculation
given as an argument and provides a rough calculation of how many
calories to consume daily in order to achieve the desired weight
loss per week.

This tool is designed to be used with input piped from a tool
such as tdee, for example:

tdee --metric --height 172 --weight 63.7 --age 29 --sex male --lifestyle 1.375 --raw |  xargs loseit --lose 0.3 --kg
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !kg && !lb {
			fmt.Println("Select kg or lb.")
			os.Exit(1)
		}

		tdee, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var deficit float64

		if kg {
			deficit = (weeklyLoss * 1000) * 1.1
		}

		if lb {
			deficit = weeklyLoss * 500
		}

		kcal := tdee - deficit
		fmt.Printf("%.0f kcal", kcal)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	weeklyLoss float64
	kg, lb     bool
)

func init() {
	RootCmd.Flags().Float64VarP(&weeklyLoss, "lose", "l", 0.0, "desired weight loss per week")
	RootCmd.Flags().BoolVar(&kg, "kg", false, "use kg for desired weight loss")
	RootCmd.Flags().BoolVar(&lb, "lb", false, "use lb for desired weight loss")
}
