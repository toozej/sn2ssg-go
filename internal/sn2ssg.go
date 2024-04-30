package sn2ssg

var (
	username          string
	password          string
	tagToDownload     string
	continuousNoteTag string
	ssgType           string
	author            string
	inputDir          string
	outputDir         string
	pollingCycle      int
	gotifyURL         string
	gotifyToken       string
	debug             bool
)

func Run() {
	getEnvVars()

	exportFromSimplenote()
}
