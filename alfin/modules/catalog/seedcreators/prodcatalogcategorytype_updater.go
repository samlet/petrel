package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProdCatalogCategoryType(ctx context.Context) error {
	log.Println("ProdCatalogCategoryType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProdCatalogCategoryType
	failures := 0

	c = cache.Get("pcct_browse_root__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Browse Root (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_browse_root__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_browse_root__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_search__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Default Search (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_search__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_search__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_other_search__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Other Search (Many)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_other_search__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_other_search__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_quick_add__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Quick Add (Many)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_quick_add__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_quick_add__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_promotions__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Promotions (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_promotions__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_promotions__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_most_popular__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Most Popular (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_most_popular__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_most_popular__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_whats_new__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("What's New (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_whats_new__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_whats_new__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_view_allw__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("View Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_view_allw__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_view_allw__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_purch_allw__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Purchase Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_purch_allw__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_purch_allw__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_admin_allw__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Admin Allow (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_admin_allw__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_admin_allw__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_ebay_root__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Ebay Root (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_ebay_root__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_ebay_root__prodcatalogcategorytype", c)
	}

	c = cache.Get("pcct_best_sell__prodcatalogcategorytype").(*ent.ProdCatalogCategoryType)
	c, err = c.Update().
		SetDescription("Best Selling (One)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcct_best_sell__prodcatalogcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcct_best_sell__prodcatalogcategorytype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
