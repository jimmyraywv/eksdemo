package cmd

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Args        int
}
