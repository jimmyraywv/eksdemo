package metrics_server

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/cmd"
	"eksdemo/pkg/helm"
	"eksdemo/pkg/template"
)

// Docs:    https://github.com/kubernetes-sigs/metrics-server/blob/master/README.md
// GitHub:  https://github.com/kubernetes-sigs/metrics-server
// Helm:    https://github.com/kubernetes-sigs/metrics-server/tree/master/charts/metrics-server
// Repo:    k8s.gcr.io/metrics-server/metrics-server
// Version: Latest is 0.5.0 (as of 10/9/21)

func NewApp() *application.Application {
	app := &application.Application{
		Command: cmd.Command{
			Name:        "metrics-server",
			Description: "Kubernetes Metric Server",
			Aliases:     []string{"metrics"},
		},

		Options: &application.ApplicationOptions{
			Namespace:      "kube-system",
			ServiceAccount: "metrics-server",
			DefaultVersion: &application.LatestPrevious{
				Latest:   "v0.5.0",
				Previous: "v0.4.4",
			},
		},

		Installer: &helm.HelmInstaller{
			ChartName:     "metrics-server",
			ReleaseName:   "metrics-server",
			RepositoryURL: "https://kubernetes-sigs.github.io/metrics-server/",
			ValuesTemplate: &template.TextTemplate{
				Template: valuesTemplate,
			},
		},
	}
	return app
}

const valuesTemplate = `
image:
  tag: {{ .Version }}
serviceAccount:
  name: {{ .ServiceAccount }}
`
