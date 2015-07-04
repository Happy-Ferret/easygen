////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"flag"
	"fmt"

	. "github.com/suntong001/easygen/easygenapi"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	// One mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		Usage()
	}
	fileName := flag.Args()[0]

	fmt.Print(Generate(Opts.HTML, fileName))
}