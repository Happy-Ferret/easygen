// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package easygen

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "easygen" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	TemplateStr  string // template string (in text)
	TemplateFile string // .tmpl (comma-separated) template file `name(s)` (default: same as .yaml file)
	ExtYaml      string // `extension` of yaml file
	ExtTmpl      string // `extension` of template file
	Debug        int    // debugging `level`
}
