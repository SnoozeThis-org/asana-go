package asana

type AccessLevel string

const (
	AccessLevelAdmin     AccessLevel = "admin"
	AccessLevelEditor    AccessLevel = "editor"
	AccessLevelCommenter AccessLevel = "commenter"
	AccessLevelViewer    AccessLevel = "viewer"
)

type ProjectMember struct {
	// Read-only. Globally unique ID of the object
	ID string `json:"gid,omitempty"`

	// Read-only. The base type of this resource
	ResourceType string `json:"resource_type,omitempty"`

	// Read-only. The name of the object.
	Name string `json:"name,omitempty"`
}

type ProjectMembership struct {
	// Read-only. Globally unique ID of the object
	ID string `json:"gid,omitempty"`

	// Read-only. The base type of this resource
	ResourceType string `json:"resource_type,omitempty"`

	// The project associated with this membership
	Parent *Project `json:"parent"`

	// Whether the member has admin, editor, commenter, or viewer access to the project.
	AccessLevel AccessLevel `json:"access_level,omitempty"`

	// Read-only. Type of the membership.
	ResourceSubtype string `json:"resource_subtype,omitempty"`
}

type membershipsRequestParams struct {
	// Globally unique identifier for goal, project, or portfolio
	Parent string `json:"parent"`
	
	// Globally unique identifier for team or user.
	Member string `json:"member"`
}

func (p *Project) Memberships(client *Client, options ...*Options) ([]*ProjectMembership,  *NextPage, error) {
	client.trace("Listing memberships in project %s...\n", p.ID)
	var result []*ProjectMembership

	// Make the request
	query := membershipsRequestParams{
		Parent: p.ID,
	}
	nextPage, err := client.get("/memberships", query, &result, options...)
	return result, nextPage, err
}
