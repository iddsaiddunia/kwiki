package tools

type Tool struct {
	Name     string
	Versions []string
	WinGet   string
	Brew     string
	Apt      string
	EnvVars  map[string]string
}

var Registry = []Tool{
	{
		Name:     "Python",
		Versions: []string{"3.13", "3.12", "3.11", "3.10"},
		WinGet:   "Python.Python.3",
		Brew:     "python@{version}",
		Apt:      "python{version}",
	},
	{
		Name:     "JDK",
		Versions: []string{"21", "17", "11"},
		WinGet:   "EclipseAdoptium.Temurin.{version}.JDK",
		Brew:     "temurin@{version}",
		Apt:      "temurin-{version}-jdk",
		EnvVars:  map[string]string{"JAVA_HOME": "auto"},
	},
	{
		Name:     "NVM",
		Versions: []string{"latest"},
		WinGet:   "CoreyButler.NVMforWindows",
		Brew:     "nvm",
		Apt:      "nvm",
	},
	{
		Name:     "Flutter",
		Versions: []string{"stable", "beta", "3.27.0", "3.24.0"},
		WinGet:   "Google.Flutter",
		Brew:     "flutter",
		Apt:      "flutter",
		EnvVars:  map[string]string{"FLUTTER_HOME": "auto"},
	},
	{
		Name:     "Android Studio",
		Versions: []string{"latest", "2024.1", "2023.3"},
		WinGet:   "Google.AndroidStudio",
		Brew:     "android-studio",
		Apt:      "android-studio",
		EnvVars:  map[string]string{"ANDROID_HOME": "auto", "ANDROID_SDK_ROOT": "auto"},
	},
	{
		Name:     "PostgreSQL",
		Versions: []string{"17", "16", "15", "14"},
		WinGet:   "PostgreSQL.PostgreSQL.{version}",
		Brew:     "postgresql@{version}",
		Apt:      "postgresql-{version}",
	},
	{
		Name:     "Git",
		Versions: []string{"latest"},
		WinGet:   "Git.Git",
		Brew:     "git",
		Apt:      "git",
	},
	{
		Name:     "VS Code",
		Versions: []string{"latest"},
		WinGet:   "Microsoft.VisualStudioCode",
		Brew:     "visual-studio-code",
		Apt:      "code",
	},
	{
		Name:     "Docker",
		Versions: []string{"latest"},
		WinGet:   "Docker.DockerDesktop",
		Brew:     "docker",
		Apt:      "docker-ce",
	},
	{
		Name:     "Node.js",
		Versions: []string{"22", "20", "18"},
		WinGet:   "OpenJS.NodeJS.LTS",
		Brew:     "node@{version}",
		Apt:      "nodejs",
	},
	{
		Name:     "Go",
		Versions: []string{"1.24", "1.23", "1.22"},
		WinGet:   "GoLang.Go",
		Brew:     "go",
		Apt:      "golang",
	},
	{
		Name:     "Rust",
		Versions: []string{"stable", "nightly"},
		WinGet:   "Rustlang.Rustup",
		Brew:     "rust",
		Apt:      "rustup",
	},
	{
		Name:     "kubectl",
		Versions: []string{"latest"},
		WinGet:   "Kubernetes.kubectl",
		Brew:     "kubectl",
		Apt:      "kubectl",
	},
	{
		Name:     "Terraform",
		Versions: []string{"latest", "1.10", "1.9"},
		WinGet:   "Hashicorp.Terraform",
		Brew:     "terraform",
		Apt:      "terraform",
	},
	{
		Name:     "Postman",
		Versions: []string{"latest"},
		WinGet:   "Postman.Postman",
		Brew:     "postman",
		Apt:      "postman",
	},
}
