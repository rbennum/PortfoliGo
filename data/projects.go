package data

type Project struct {
	Name       string
	Desc       string
	ProjectURL string
	GithubURL  string
	Tech       []string
}

var Projects = []Project{
	{
		Name:       "PortfoliGo",
		Desc:       "My portfolio page. You're looking at it lol.",
		ProjectURL: "#",
		GithubURL:  "https://github.com/rbennum/PortfoliGo",
		Tech:       []string{"HTML", "CSS", "Golang", "Docker"},
	},
	{
		Name:       "url-shrtnr",
		Desc:       "A web app to shorten your URL.",
		ProjectURL: "https://rbennum.bid/shortener/",
		GithubURL:  "https://github.com/rbennum/url-shrtnr",
		Tech: []string{
			"HTML",
			"CSS",
			"JS",
			"Golang",
			"PostgreSQL",
			"Docker",
		},
	},
	{
		Name:       "Alfagift",
		Desc:       "An e-commerce made in iOS app, helped to develop and maintain its features back when I was working for Alfagift.",
		ProjectURL: "https://apps.apple.com/id/app/alfagift-alfamart-online-shop/id1013717463",
		GithubURL:  "",
		Tech:       []string{"Swift", "UIKit", "Storyboard"},
	},
	{
		Name:       "Ajaib",
		Desc:       "An iOS app to sell financial products such as stocks and mutual funds. I developed for the Authentication and Onboarding modules, making their app safer when I was working for Ajaib.",
		ProjectURL: "https://apps.apple.com/id/app/ajaib-saham-bond-reksadana/id1473916571",
		GithubURL:  "",
		Tech:       []string{"Swift", "UIKit", "SnapKit", "TCA"},
	},
	{
		Name:       "Ajaib Kripto",
		Desc:       "Imagine the previous app, but for selling crypto stuff. Also developing the same modules, as it was on the same repo. So it's like two sides of the same coin (am I using it right?).",
		ProjectURL: "",
		GithubURL:  "https://apps.apple.com/id/app/ajaib-kripto-buy-btc-crypto/id1634168301",
		Tech:       []string{"Swift", "UIKit", "SnapKit", "TCA"},
	},
}

type Tech struct {
	Name     string
	ImageURL string
}

var TechList = []Tech{
	{Name: "Swift", ImageURL: "/dist/assets/swift.svg"},
	{Name: "Flutter", ImageURL: "/dist/assets/flutter.svg"},
	{Name: "Golang", ImageURL: "/dist/assets/golang.svg"},
	{Name: "PostgreSQL", ImageURL: "/dist/assets/postgresql.svg"},
	{Name: "Docker", ImageURL: "/dist/assets/docker.svg"},
}

type PageData struct {
	ProjectList []Project
	TechList    []Tech
}

var PageDataItem = PageData{
	ProjectList: Projects,
	TechList:    TechList,
}
