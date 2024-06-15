package main

import (
	"log"
	"os"
	"path"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Resume struct {
	Basics          Basics        `yaml:"basics"`
	Accomplishments []string      `yaml:"accomplishments"`
	Works           []Work        `yaml:"work"`
	Volunteers      []Volunteer   `yaml:"volunteer"`
	Educations      []Education   `yaml:"education"`
	Awards          []Award       `yaml:"awards"`
	Certificates    []Certificate `yaml:"certificates"`
	Publications    []Publication `yaml:"publications"`
	Skills          []Skill       `yaml:"skills"`
	Languages       []Language    `yaml:"languages"`
	Interests       []Interest    `yaml:"interests"`
	References      []Reference   `yaml:"references"`
	Projects        []Project     `yaml:"projects"`
}

type Basics struct {
	Name     string    `yaml:"name"`
	Label    string    `yaml:"label"`
	Image    string    `yaml:"image"`
	Email    string    `yaml:"email"`
	Phone    string    `yaml:"phone"`
	Url      string    `yaml:"url"`
	Summary  string    `yaml:"summary"`
	Location Location  `yaml:"location"`
	Profiles []Profile `yaml:"profiles"`
}

type Location struct {
	Address     string `yaml:"address"`
	PostalCode  string `yaml:"postalCode"`
	City        string `yaml:"city"`
	CountryCode string `yaml:"countryCode"`
	Region      string `yaml:"region"`
}

type Profile struct {
	Network  string `yaml:"network"`
	Username string `yaml:"username"`
	URL      string `yaml:"url"`
}

type Work struct {
	Name       string    `yaml:"name"`
	Position   string    `yaml:"position"`
	URL        string    `yaml:"url"`
	StartDate  time.Time `yaml:"startDate"`
	EndDate    time.Time `yaml:"endDate"`
	Summary    string    `yaml:"summary"`
	Highlights []string  `yaml:"highlights"`
}

type Volunteer struct {
	Organization string    `yaml:"organization"`
	Position     string    `yaml:"position"`
	URL          string    `yaml:"url"`
	StartDate    time.Time `yaml:"startDate"`
	EndDate      time.Time `yaml:"endDate"`
	Summary      string    `yaml:"summary"`
	Highlights   []string  `yaml:"highlights"`
}

type Education struct {
	Institution    string    `yaml:"institution"`
	URL            string    `yaml:"url"`
	Area           string    `yaml:"area"`
	StudyType      string    `yaml:"studyType"`
	GraduationDate time.Time `yaml:"graduationDate"`
	Score          string    `yaml:"score"`
	Courses        []string  `yaml:"courses"`
}

type Award struct {
	Title   string    `yaml:"title"`
	Date    time.Time `yaml:"date"`
	Awarder string    `yaml:"awarder"`
	Summary string    `yaml:"summary"`
}

type Certificate struct {
	Name          string    `yaml:"name"`
	IssueDate     time.Time `yaml:"issueDate"`
	ExpireDate    time.Time `yaml:"expireDate"`
	CertificateID string    `yaml:"certificateID"`
	Issuer        string    `yaml:"issuer"`
	URL           string    `yaml:"url"`
}

type Publication struct {
	Name        string    `yaml:"name"`
	Publisher   string    `yaml:"publisher"`
	ReleaseDate time.Time `yaml:"releaseDate"`
	URL         string    `yaml:"url"`
	Summary     string    `yaml:"summary"`
}

type Skill struct {
	Name     string   `yaml:"name"`
	Level    string   `yaml:"level"`
	Keywords []string `yaml:"keywords"`
}

type Language struct {
	Language string `yaml:"language"`
	Fluency  string `yaml:"fluency"`
}

type Interest struct {
	Name     string   `yaml:"name"`
	Keywords []string `yaml:"keywords"`
}

type Reference struct {
	Name      string `yaml:"name"`
	Reference string `yaml:"reference"`
}

type Project struct {
	Name        string    `yaml:"name"`
	StartDate   time.Time `yaml:"startDate"`
	EndDate     time.Time `yaml:"endDate"`
	Description string    `yaml:"description"`
	Highlights  []string  `yaml:"highlights"`
	URL         string    `yaml:"url"`
}

func main() {
	app := &cli.App{
		Name:      "yamlresume",
		Usage:     "convert yaml resume to html",
		ArgsUsage: "<input yaml> <input template> <output filename>",
		Action: func(cCtx *cli.Context) error {
			var inFile = cCtx.Args().Get(0)
			var tmplFile = cCtx.Args().Get(1)
			var outHtml = cCtx.Args().Get(2)

			data, err := os.ReadFile(inFile)
			if err != nil {
				panic(err)
			}

			funcMap := template.FuncMap{
				"inc": func(i int) int {
					return i + 1
				},
			}

			t, err := template.New(path.Base(tmplFile)).Funcs(funcMap).ParseFiles(tmplFile)
			if err != nil {
				panic(err)
			}

			var resume Resume

			if err := yaml.Unmarshal(data, &resume); err != nil {
				panic(err)
			}

			var f *os.File
			f, err = os.Create(outHtml)
			if err != nil {
				panic(err)
			}

			err = t.Execute(f, resume)
			if err != nil {
				panic(err)
			}

			err = f.Close()
			if err != nil {
				panic(err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
