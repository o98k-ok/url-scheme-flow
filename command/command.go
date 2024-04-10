package command

type Commander interface {
	GenURI() string
	IconApp() string
	Filtered(keys []string) (string, string, bool)
	SetArgs(v string)
	Title() (string, string)
}
