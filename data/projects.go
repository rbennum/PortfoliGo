package data

type Project struct {
	Name		string
	Description	string
	URL			string
}

var Projects = []Project{
	{ 
		Name: "Project 1",
		Description: "Basic description",
		URL: "/project-1", 
	},
}