package main

import (
	namegen "github.com/BenjaminNolan/go-namegen"

	"github.com/spf13/cobra"
)

var (
	capitalize bool
	alliterate bool
	separator  string
)

func main() {
	cmd := &cobra.Command{
		Use:   "generate-name",
		Short: "Generate a random name",
		Run: func(cmd *cobra.Command, args []string) {
			generateName()
		},
	}

	cmd.Flags().BoolVarP(&capitalize, "capitalize", "c", true, "Capitalize the first letter of words")
	cmd.Flags().BoolVarP(&alliterate, "alliterate", "a", false, "Generate alliterative names, eg, 'Awesome Antelope' or 'Clever Cat'")
	cmd.Flags().StringVarP(&separator, "separator", "s", " ", "The separator between words, eg, '-', ' ', or '_'")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func generateName() {
	generator := namegen.New()

	generator.Capitalize = capitalize
	generator.Alliterate = alliterate
	generator.Separator = separator

	println(generator.Generate())
}
