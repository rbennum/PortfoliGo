package data

type Project struct {
	Name      string
	ShortDesc string
	LongDesc  string
	URL       string
}

var Projects = []Project{
	{
		Name:      "Project 1",
		ShortDesc: "Basic description",
		LongDesc:  "Long basic description 2 paragraphs minimum lol",
		URL:       "/project-1",
	},
}
