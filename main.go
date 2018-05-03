/*
GcodeDetails

Description: An executable to display details of the indicated GCODE file,
						 suitable for understanding these toolpath files.

Notes:       The current version expects Cura as the slicer and has been tested
						 with v2.3.1 of same.

Author:      Michael Blankenship
Repo:        https://github.com/OutsourcedGuru/GcodeDetails
*/
package main
import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func syntax() {
	fmt.Printf("Syntax:  GcodeDetails GCodeFilePath\n\n")
	os.Exit(1)
}

func categorizeLine(input string) (output string) {
	if input == "" {
		return ""
	}
	switch []rune(input)[0] {
	case 'G':
		return "Standard GCode command, such as move to a point"
	case 'M':
		return "RepRap-defined command, such as turn on a cooling fan"
	case 'T':
		return "Select tool nnn. In RepRap, a tool is typically associated with a nozzle, which may be fed by one or more extruders."
	case 'S':
		return "Command parameter, such as time in seconds; temperatures; voltage to send to a motor"
	case 'P':
		return "Command parameter, such as time in milliseconds; proportional (Kp) in PID Tuning"
	case 'X':
		return "A X coordinate, usually to move to. This can be an Integer or Fractional number."
	case 'Y':
		return "A Y coordinate, usually to move to. This can be an Integer or Fractional number."
	case 'Z':
		return "A Z coordinate, usually to move to. This can be an Integer or Fractional number."
	case 'U':
		return "Additional axis coordinates (RepRapFirmware)"
	case 'V':
		return "Additional axis coordinates (RepRapFirmware)"
	case 'W':
		return "Additional axis coordinates (RepRapFirmware)"
	case 'I':
		return "Parameter - X-offset in arc move; integral (Ki) in PID Tuning"
	case 'J':
		return "Parameter - Y-offset in arc move"
	case 'D':
		return "Parameter - used for diameter; derivative (Kd) in PID Tuning"
	case 'H':
		return "Parameter - used for heater number in PID Tuning"
	case 'F':
		return "Feedrate in mm per minute. (Speed of print head movement)"
	case 'R':
		return "Parameter - used for temperatures"
	case 'Q':
		return "Parameter - not currently used"
	case 'E':
		return "Length of extrudate. This is exactly like X, Y and Z, but for the length of filament to consume."
	case 'N':
		return "Line number. Used to request repeat transmission in the case of communications errors."
	case '*':
		return "Checksum. Used to check for communications errors."
	case ';':
		return "Comment"
	default:
		return "?"
	}


}

func main() {
	bReadError                      := false
	inputfilename                   := "N/A"
	flag.Parse()
	if len(flag.Args()) != 1        { syntax() }		// Should be only one argument as a filename
	
	inputfilename = flag.Args()[0]
	data, err := ioutil.ReadFile(inputfilename)
	if err != nil {
		bReadError = true;
		fmt.Fprintf(os.Stderr, "GcodeDetails:\n  %v\n\n", err)
		return
	}

	// Now process the input file into the output file
	for _, line := range strings.Split(string(data), "\n") {
		output := categorizeLine(line)
		fmt.Printf("Read: [%s][%s]\n", line, output)
	}
	if !bReadError {
		fmt.Printf("\n\n\nInput:  %s\n", inputfilename)
		fmt.Printf("\nFinished.\n\n")
	}
}