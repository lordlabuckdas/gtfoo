package gtfobins

const (
	colorReset      = "\033[0m"
	colorRed        = "\033[91m"
	colorGreen      = "\033[92m"
	colorYellow     = "\033[93m"
	colorBlue       = "\033[94m"
	colorPurple     = "\033[95m"
	colorCyan       = "\033[96m"
	colorWhite      = "\033[97m"
	shell           = "Shell"
	cmd             = "Command"
	revShell        = "Reverse Shell"
	nonIntRevShell  = "Non-Interactive Reverse Shell"
	bindShell       = "Bind Shell"
	nonIntBindShell = "Non-Interactive Bind Shell"
	fileUpload      = "File Upload"
	fileDownload    = "File Download"
	fileWrite       = "File Write"
	fileRead        = "File Read"
	libLoad         = "Library Load"
	suid            = "SUID"
	sudo            = "Sudo"
	capab           = "Capabilities"
	limSUID         = "Limited SUID"
)

var baseURL = "https://raw.githubusercontent.com/GTFOBins/GTFOBins.github.io/master/_gtfobins/"

type function struct {
	Description string `yaml:"description,omitempty"`
	Code        string `yaml:"code,omitempty"`
}

type gtfoStruct struct {
	Functions struct {
		Shell                      []function `yaml:"shell"`
		Command                    []function `yaml:"command"`
		ReverseShell               []function `yaml:"reverse-shell"`
		NonInteractiveReverseShell []function `yaml:"non-interactive-reverse-shell"`
		BindShell                  []function `yaml:"bind-shell"`
		NonInteractiveBindShell    []function `yaml:"non-interactive-bind-shell"`
		FileUpload                 []function `yaml:"file-upload"`
		FileDownload               []function `yaml:"file-download"`
		FileWrite                  []function `yaml:"file-write"`
		FileRead                   []function `yaml:"file-read"`
		LibraryLoad                []function `yaml:"library-load"`
		SUID                       []function `yaml:"suid"`
		Sudo                       []function `yaml:"sudo"`
		Capabilities               []function `yaml:"capabilities"`
		LimitedSuid                []function `yaml:"limited-suid"`
	} `yaml:"functions"`
}

var funcDesc = map[string]string{
	"Shell":                         "It can be used to break out from restricted environments by spawning an interactive system shell.",
	"Command":                       "It can be used to break out from restricted environments by running non-interactive system commands.",
	"Reverse Shell":                 "It can send back a reverse shell to a listening attacker to open a remote network access.",
	"Non-Interactive Reverse Shell": "It can send back a non-interactive reverse shell to a listening attacker to open a remote network access.",
	"Bind Shell":                    "It can bind a shell to a local port to allow remote network access.",
	"Non-Interactive Bind Shell":    "It can bind a non-interactive shell to a local port to allow remote network access.",
	"File Upload":                   "It can exfiltrate files on the network.",
	"File Download":                 "It can download remote files.",
	"File Write":                    "It writes data to files, it may be used to do privileged writes or write files outside a restricted file system.",
	"File Read":                     "It reads data from files, it may be used to do privileged reads or disclose files outside a restricted file system.",
	"Library Load":                  "It loads shared libraries that may be used to run code in the binary execution context.",
	"SUID":                          "If the binary has the SUID bit set, it does not drop the elevated privileges and may be exploited to access the file system, escalate or maintain privileged access as a SUID backdoor. If it is used to run `sh -p`, omit the `-p` argument on systems like Debian (<= Stretch) that allow the default `sh` shell to run with SUID privileges.\nThis example creates a local SUID copy of the binary and runs it to maintain elevated privileges. To exploit an existing SUID binary skip the first command and run the program using its original path.",
	"Sudo":                          "If the binary is allowed to run as superuser by `sudo`, it does not drop the elevated privileges and may be used to access the file system, escalate or maintain privileged access.",
	"Capabilities":                  "If the binary has the Linux `CAP_SETUID` capability set or it is executed by another binary with the capability set, it can be used as a backdoor to maintain privileged access by manipulating its own process UID.",
	"Limited SUID":                  "If the binary has the SUID bit set, it may be exploited to access the file system, escalate or maintain access with elevated privileges working as a SUID backdoor. If it is used to run commands (e.g., via `system()`-like invocations) it only works on systems like Debian (<= Stretch) that allow the default `sh` shell to run with SUID privileges.\nThis example creates a local SUID copy of the binary and runs it to maintain elevated privileges. To exploit an existing SUID binary skip the first command and run the program using its original path.",
}
