package asset

import (
	"context"
	"fmt"
	"github.com/samlet/petrel/alfin/modules/asset/ent"
	"github.com/samlet/petrel/alfin/modules/asset/ent/asset"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestAsset(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func IdRef(ctx context.Context, client *ent.Client, stringId string) *ent.Asset {
	rec, err := client.Asset.
		Query().
		Where(asset.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("returned: ", rec)
	return rec
}

func Do(ctx context.Context, client *ent.Client) error {
	asset, err:=client.Asset.Create().SetModel("entity").
		SetRegisteredAt(time.Now()).
		SetStringRef("first").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating asset: %w", err)
	}
	pkg, err:= client.WorkloadPkg.
		Create().
		SetName("cadence").
		AddAssets(asset).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pkg: %w", err)
	}
	log.Println("pkg was created: ", pkg)

	// query
	pkgCadence:=client.WorkloadPkg.Query().Where(workloadpkg.Name("cadence")).OnlyX(ctx)
	log.Println("found: ", pkgCadence)
	assets, err:=pkgCadence.QueryAssets().All(ctx)
	if err != nil {
		return fmt.Errorf("query assets: %w", err)
	}
	log.Println("assets: ", assets)

	// query by string-ref
	println("id-ref for first: ", IdRef(ctx, client, "first").ID)
	return nil
}
