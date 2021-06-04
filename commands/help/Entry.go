package help

import (
	"internal/general"
	"pkg/color"
)

// Entry function is the entry point for
// the devices module.
func Entry(args []string) {

	// Get latest commands from server
	var c Commands
	c.Fill()

	general.PrintHeading("Welcome to Lux!", color.FgWhite)
	general.PrintHeading("Lux " + general.BuildName, color.Italic)
	general.Line()
	c.List()
	general.Line()
	general.PrintHeading("ABOUT LUX", color.FgWhite)
	general.PrintStringParagraph("description", general.Description, color.FgWhite)
	general.PrintStringParagraph("build", general.BuildName, color.FgWhite)
	general.PrintStringParagraph("license", general.License, color.FgWhite)
	general.PrintStringParagraph("repository", general.GHRepo, color.FgWhite)
	general.PrintBoolParagraph("up to date", color.FgWhite, general.BuildNo == c.BuildNo)
}
