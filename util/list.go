package util

import (
	"fmt"

	"bitbucket.org/mikehouston/asana-go"
)

func ListWorkspaces(c *asana.Client) error {
	// List workspaces
	workspaces, nextPage, err := c.Workspaces()
	if err != nil {
		return err
	}
	_ = nextPage

	for _, workspace := range workspaces {
		if workspace.IsOrganization {
			fmt.Printf("Organization %s %q\n", workspace.ID, workspace.Name)
		} else {
			fmt.Printf("Workspace %s %q\n", workspace.ID, workspace.Name)
		}
	}
	return nil
}

func ListProjects(client *asana.Client, w *asana.Workspace) error {
	// List projects
	projects, nextPage, err := w.Projects(client)
	if err != nil {
		return err
	}
	_ = nextPage

	for _, project := range projects {
		fmt.Printf("Project %s: %q\n", project.ID, project.Name)
	}
	return nil
}

func ListTasks(client *asana.Client, p *asana.Project) error {
	// List projects
	tasks, nextPage, err := p.Tasks(client)
	if err != nil {
		return err
	}
	_ = nextPage

	for _, task := range tasks {
		fmt.Printf("Task %s %q\n", task.ID, task.Name)
	}
	return nil
}

func ListSections(client *asana.Client, p *asana.Project) error {
	// List sections
	sections, nextPage, err := p.Sections(client)
	if err != nil {
		return err
	}
	_ = nextPage

	for _, section := range sections {
		fmt.Printf("Section %s %q\n", section.ID, section.Name)
	}
	return nil
}
