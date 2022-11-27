package attendance

/*
func loadTableOfOrganizationConfiguration(toCfg string) (pkg.TableOfOrganization, error) {
	daCfg := yaml.DutyAssignmentConfig{}
	err := yaml.LoadYamlDocFromFile(toCfg, &daCfg)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainDACfg := daCfg.DomainDutyAssignments()

	to := yaml.TableOfOrganization{}
	err = yaml.LoadYamlDocFromFile(filepath.Join(testDataDir, "to.yaml"), &to)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	domainTO, err := to.DomainTableOfOrganization(domainDACfg)
	if err != nil {
		return pkg.TableOfOrganization{}, err
	}

	return domainTO, nil
}
*/

func loadCapwatchData() error {
	return nil
}

func generateLaTeX() error {
	return nil
}

func compileLaTeX() error {
	return nil
}

func moveCompiledLog() error {
	return nil
}

func BuildBarcodeLog(input, output, password string) error {
	// Load TO
	// Load CAPWATCH data
	// Generate LaTeX
	// Compile LaTeX

	return nil
}
