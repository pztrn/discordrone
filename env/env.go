package env

import (
	// stdlib
	"fmt"

	// other
	"github.com/vrischmann/envconfig"
)

var Data *environmentData

type environmentData struct {
	CI      bool `envconfig:"optional"`
	IsDrone bool `envconfig:"DRONE,optional"`
	Debug   bool `envconfig:"optional"`

	Drone struct {
		Branch string `envconfig:"optional"`
		Build  struct {
			Action      string  `envconfig:"optional"`
			Created     float64 `envconfig:"optional"`
			CreatedInt  int     `envconfig:"optional"`
			Event       string  `envconfig:"optional"`
			Finished    float64 `envconfig:"optional"`
			FinishedInt int     `envconfig:"optional"`
			Link        string  `envconfig:"optional"`
			Number      int     `envconfig:"optional"`
			Parent      string  `envconfig:"optional"`
			Started     float64 `envconfig:"optional"`
			StartedInt  int     `envconfig:"optional"`
			Status      string  `envconfig:"optional"`
		}
		CommitHash string `envconfig:"DRONE_COMMIT,optional"`
		Commit     struct {
			After      string `envconfig:"optional"`
			AuthorName string `envconfig:"DRONE_COMMIT_AUTHOR,optional"`
			Author     struct {
				Avatar string `envconfig:"optional"`
				Email  string `envconfig:"optional"`
				Name   string `envconfig:"optional"`
			}
			Before  string `envconfig:"optional"`
			Branch  string `envconfig:"optional"`
			Link    string `envconfig:"optional"`
			Message string `envconfig:"optional"`
			Ref     string `envconfig:"optional"`
			Sha     string `envconfig:"optional"`
		}
		Deploy struct {
			To string `envconfig:"optional"`
		}
		Failed struct {
			Stages string `envconfig:"optional"`
			Steps  string `envconfig:"optional"`
		}
		Git struct {
			HTTP struct {
				URL string `envconfig:"optional"`
			}
			SSH struct {
				URL string `envconfig:"optional"`
			}
		}
		Pull struct {
			Request string `envconfig:"optional"`
		}
		Remote struct {
			URL string `envconfig:"optional"`
		}
		RepoName string `envconfig:"DRONE_REPO_NAME,optional"`
		Repo     struct {
			Branch     string `envconfig:"optional"`
			Link       string `envconfig:"optional"`
			Name       string `envconfig:"optional"`
			Namespace  string `envconfig:"optional"`
			Owner      string `envconfig:"optional"`
			Private    bool   `envconfig:"optional"`
			SCM        string `envconfig:"optional"`
			Visibility string `envconfig:"optional"`
		}
		SemverData string `envconfig:"DRONE_SEMVER,optional"`
		Semver     struct {
			Build      string `envconfig:"optional"`
			Error      string `envconfig:"optional"`
			Major      int    `envconfig:"optional"`
			Minor      int    `envconfig:"optional"`
			Patch      int    `envconfig:"optional"`
			Prerelease string `envconfig:"optional"`
			Short      string `envconfig:"optional"`
		}
		Source struct {
			Branch string `envconfig:"optional"`
		}
		Stage struct {
			Arch    string `envconfig:"optional"`
			Depends struct {
				On string `envconfig:"optional"`
			}
			Finished int    `envconfig:"optional"`
			Kind     string `envconfig:"optional"`
			Machine  string `envconfig:"optional"`
			Name     string `envconfig:"optional"`
			Number   int    `envconfig:"optional"`
			OS       string `envconfig:"optional"`
			Started  int    `envconfig:"optional"`
			Status   string `envconfig:"optional"`
			Type     string `envconfig:"optional"`
			Variant  string `envconfig:"optional"`
		}
		Step struct {
			Name   string `envconfig:"optional"`
			Number int    `envconfig:"optional"`
		}
		System struct {
			Host     string `envconfig:"optional"`
			Hostname string `envconfig:"optional"`
			Proto    string `envconfig:"optional"`
			Version  string `envconfig:"optional"`
		}
		Tag    string `envconfig:"optional"`
		Target struct {
			Branch string `envconfig:"optional"`
		}
	}
	Plugin struct {
		Avatar struct {
			URL string `envconfig:"optional"`
		}
		Color   string `envconfig:"optional"`
		Message string `envconfig:"optional"`
		Webhook struct {
			ID    string `envconfig:"optional"`
			Token string `envconfig:"optional"`
		}
		TTS      bool   `envconfig:"optional"`
		Username string `envconfig:"optional"`
	}
}

func ParseEnv() {
	Data = &environmentData{}

	err := envconfig.Init(Data)
	if err != nil {
		fmt.Println(err)
	}
}

func (ed *environmentData) ConvertFloatToInts() {
	ed.Drone.Build.CreatedInt = int(ed.Drone.Build.Created)
	ed.Drone.Build.FinishedInt = int(ed.Drone.Build.Finished)
	ed.Drone.Build.StartedInt = int(ed.Drone.Build.Started)
}
