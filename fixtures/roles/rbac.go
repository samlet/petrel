package roles

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

type Security struct{
	e *casbin.Enforcer
}

func NewSecurity() (*Security, error) {
	// Initialize a Xorm adapter with MySQL database.
	a, err := xormadapter.NewAdapter("mysql",
		"root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
		return nil,err
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
		return nil,err
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
		return nil,err
	}
	return &Security{e: e}, nil
}

// Check the security
// 	sub := "alice" // the user that wants to access a resource.
//	obj := "data1" // the resource that is going to be accessed.
//	act := "read" // the operation that the user performs on the resource.
func (t Security) Check(sub, obj, act string) (bool, error) {
	return t.e.Enforce(sub, obj, act)
}

