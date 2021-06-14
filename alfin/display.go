package alfin

import (
	log "github.com/sirupsen/logrus"
	"github.com/xlab/treeprint"
	"strings"
)

func DisplayEntInfo(ent *ModelEntity, tree treeprint.Tree, showFields bool) {
	type ChildrenRel struct {
		rel, name string
		keyMaps   []ModelKeymap
	}
	entNode := tree.AddMetaBranch("entity", ent.Name)
	children := []ChildrenRel{}
	for _, edge := range ent.Edges() {
		switch edge.Type {
		case "many":
			//entNode.AddMetaBranch("*"+edge.Name, edge.RelEntityName)
			children = append(children, ChildrenRel{
				edge.Name,
				edge.RelEntityName,
				edge.Keymaps,
			})
		case "one-nofk":
			entNode.AddMetaBranch("-"+edge.Name, edge.RelEntityName)
		case "one":
			br:=entNode.AddMetaBranch("+"+edge.Name, edge.RelEntityName)
			for _, cr := range edge.Keymaps {
				br.AddMetaNode(cr.FieldName, cr.RelFieldName)
			}
		default:
			log.Fatal("Unknown edge type", edge.Type)
		}
	}

	if len(children) == 0 {
		entNode.AddNode("no children")
	} else {
		childrenNodes := entNode.AddBranch("children")
		for _, c := range children {
			br:=childrenNodes.AddMetaBranch(c.rel, c.name)
			for _, cr := range c.keyMaps {
				br.AddMetaNode(cr.FieldName, cr.RelFieldName)
			}
		}
	}

	if showFields {
		fieldNodes := entNode.AddBranch("fields")
		fieldNodes.AddMetaNode("pk", strings.Join(ent.Pks, ", "))
		for _, fld := range ent.NormalFields() {
			var fldMeta string
			if fld.Pk{
				fldMeta="pk"
			}else{
				fldMeta=fld.Type
			}
			fieldNodes.AddMetaNode(fldMeta, fld.Name)
		}
	}
}
