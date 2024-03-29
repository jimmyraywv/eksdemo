package kubernetes

import (
	"eksdemo/pkg/resource"
	"eksdemo/pkg/template"
	"fmt"
)

type ResourceManager struct {
	template.Template
	DryRun bool
}

func (m *ResourceManager) Create(options resource.Options) error {
	manifest, err := m.Render(options)
	if err != nil {
		return err
	}

	if m.DryRun {
		fmt.Println("\nKubernetes Resource Manager Dry Run:")
		fmt.Println(manifest)
		return nil
	}

	return CreateResources(options.GetKubeContext(), manifest)
}

func (m *ResourceManager) Delete(options resource.Options) error {
	return nil
}

func (m *ResourceManager) SetDryRun() {
	m.DryRun = true
}
