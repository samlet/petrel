package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProdCatalogCategoryType(ctx context.Context) error {
	log.Println("ProdCatalogCategoryType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProdCatalogCategoryType

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_browse_root__prodcatalogcategorytype").
		SetDescription("Browse Root (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_browse_root__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_browse_root__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_search__prodcatalogcategorytype").
		SetDescription("Default Search (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_search__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_search__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_other_search__prodcatalogcategorytype").
		SetDescription("Other Search (Many)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_other_search__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_other_search__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_quick_add__prodcatalogcategorytype").
		SetDescription("Quick Add (Many)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_quick_add__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_quick_add__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_promotions__prodcatalogcategorytype").
		SetDescription("Promotions (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_promotions__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_promotions__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_most_popular__prodcatalogcategorytype").
		SetDescription("Most Popular (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_most_popular__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_most_popular__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_whats_new__prodcatalogcategorytype").
		SetDescription("What's New (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_whats_new__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_whats_new__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_view_allw__prodcatalogcategorytype").
		SetDescription("View Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_view_allw__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_view_allw__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_purch_allw__prodcatalogcategorytype").
		SetDescription("Purchase Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_purch_allw__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_purch_allw__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_admin_allw__prodcatalogcategorytype").
		SetDescription("Admin Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_admin_allw__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_admin_allw__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_ebay_root__prodcatalogcategorytype").
		SetDescription("Ebay Root (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_ebay_root__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_ebay_root__prodcatalogcategorytype", c)

	c, err = client.ProdCatalogCategoryType.Create().SetStringRef("pcct_best_sell__prodcatalogcategorytype").
		SetDescription("Best Selling (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcct_best_sell__prodcatalogcategorytype: %v", err)
		return err
	}
	cache.Put("pcct_best_sell__prodcatalogcategorytype", c)

	return nil
}
