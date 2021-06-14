package alfin

import (
	log "github.com/sirupsen/logrus"
	"github.com/xlab/treeprint"
)

func DisplayEntInfo(ent *ModelEntity, tree treeprint.Tree) {
	type ChildrenRel struct{ rel, name string }
	entNode := tree.AddMetaBranch("entity", ent.Name)
	children := []ChildrenRel{}
	for _, edge := range ent.Edges() {
		switch edge.Type {
		case "many":
			//entNode.AddMetaBranch("*"+edge.Name, edge.RelEntityName)
			children = append(children, ChildrenRel{edge.Name, edge.RelEntityName})
		case "one-nofk":
			entNode.AddMetaBranch("-"+edge.Name, edge.RelEntityName)
		case "one":
			entNode.AddMetaBranch("+"+edge.Name, edge.RelEntityName)
		default:
			log.Fatal("Unknown edge type", edge.Type)
		}

	}

	childrenNodes := entNode.AddBranch("children")
	for _, c := range children {
		childrenNodes.AddMetaNode(c.rel, c.name)
	}
}

