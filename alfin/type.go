package alfin

type (
	ModelEntity struct{
		Name string `json:"name"`
		Fields []*ModelField `json:"fields"`
	}

	ModelField struct {
		Name string `json:"name"`
	}
)

