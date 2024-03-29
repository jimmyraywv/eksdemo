package helm

import (
	"eksdemo/pkg/application"
	"eksdemo/pkg/template"
	"fmt"
)

type HelmInstaller struct {
	ChartName      string
	DryRun         bool
	ReleaseName    string
	RepositoryURL  string
	ValuesTemplate template.Template
	VersionField   string
	Wait           bool
}

func (h *HelmInstaller) Install(options application.Options) error {

	valuesFile, err := h.ValuesTemplate.Render(options)
	if err != nil {
		return err
	}

	ic := &InstallConfiguration{
		AppVersion:    options.Common().Version,
		ChartName:     h.ChartName,
		Namespace:     options.Common().Namespace,
		ReleaseName:   h.ReleaseName,
		RepositoryURL: h.RepositoryURL,
		ValuesFile:    valuesFile,
		Wait:          h.Wait,
	}

	if h.DryRun {
		fmt.Println("\nHelm Installer Dry Run:")
		fmt.Printf("%+v\n", ic)
		return nil
	}

	return Install(ic, options.KubeContext())
}

func (h *HelmInstaller) SetDryRun() {
	h.DryRun = true
}

func (h *HelmInstaller) Uninstall(options application.Options) error {
	o := options.Common()

	fmt.Printf("Checking status of Helm release: %s, in namespace: %s\n", h.ReleaseName, o.Namespace)
	if _, err := Status(o.KubeContext(), h.ReleaseName, o.Namespace); err != nil {
		return err
	}

	fmt.Println("Status validated. Uninstalling...")
	return Uninstall(o.KubeContext(), h.ReleaseName, o.Namespace)
}
