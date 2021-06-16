package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductContentType(ctx context.Context) error {
	log.Println("ProductContentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductContentType
	failures := 0

	c = cache.Get("online_access__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Online Access").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update online_access__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("online_access__productcontenttype", c)
	}

	c = cache.Get("digital_download__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Digital Download").
		AddChildren(cache.Get("installation__productcontenttype").(*ent.ProductContentType)).
		AddChildren(cache.Get("specification__productcontenttype").(*ent.ProductContentType)).
		AddChildren(cache.Get("warranty__productcontenttype").(*ent.ProductContentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update digital_download__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("digital_download__productcontenttype", c)
	}

	c = cache.Get("fulfillment_email__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Fulfillment Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fulfillment_email__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fulfillment_email__productcontenttype", c)
	}

	c = cache.Get("fulfillment_extasync__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Fulfillment External (Async)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fulfillment_extasync__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fulfillment_extasync__productcontenttype", c)
	}

	c = cache.Get("fulfillment_extsync__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Fulfillment External (Sync)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fulfillment_extsync__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fulfillment_extsync__productcontenttype", c)
	}

	c = cache.Get("product_name__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Name").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_name__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_name__productcontenttype", c)
	}

	c = cache.Get("description__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update description__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("description__productcontenttype", c)
	}

	c = cache.Get("long_description__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update long_description__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("long_description__productcontenttype", c)
	}

	c = cache.Get("alternative_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Alternative URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update alternative_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("alternative_url__productcontenttype", c)
	}

	c = cache.Get("price_detail_text__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Price Detail Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_detail_text__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_detail_text__productcontenttype", c)
	}

	c = cache.Get("ingredients__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ingredients").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ingredients__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ingredients__productcontenttype", c)
	}

	c = cache.Get("unique_ingredients__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Unique Ingredients").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update unique_ingredients__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("unique_ingredients__productcontenttype", c)
	}

	c = cache.Get("specialinstructions__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Special Instructions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update specialinstructions__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("specialinstructions__productcontenttype", c)
	}

	c = cache.Get("warnings__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Warnings").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update warnings__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("warnings__productcontenttype", c)
	}

	c = cache.Get("directions__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Directions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update directions__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("directions__productcontenttype", c)
	}

	c = cache.Get("terms_and_conds__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Terms and Conditions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update terms_and_conds__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("terms_and_conds__productcontenttype", c)
	}

	c = cache.Get("delivery_info__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Delivery Info").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update delivery_info__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("delivery_info__productcontenttype", c)
	}

	c = cache.Get("small_image_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update small_image_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("small_image_url__productcontenttype", c)
	}

	c = cache.Get("medium_image_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update medium_image_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("medium_image_url__productcontenttype", c)
	}

	c = cache.Get("large_image_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update large_image_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("large_image_url__productcontenttype", c)
	}

	c = cache.Get("detail_image_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update detail_image_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("detail_image_url__productcontenttype", c)
	}

	c = cache.Get("original_image_url__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Original").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update original_image_url__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("original_image_url__productcontenttype", c)
	}

	c = cache.Get("small_image_alt__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Alt Text - Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update small_image_alt__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("small_image_alt__productcontenttype", c)
	}

	c = cache.Get("medium_image_alt__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Alt Text - Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update medium_image_alt__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("medium_image_alt__productcontenttype", c)
	}

	c = cache.Get("large_image_alt__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Alt Text - Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update large_image_alt__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("large_image_alt__productcontenttype", c)
	}

	c = cache.Get("detail_image_alt__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Alt Text - Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update detail_image_alt__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("detail_image_alt__productcontenttype", c)
	}

	c = cache.Get("additional_image_1__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update additional_image_1__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("additional_image_1__productcontenttype", c)
	}

	c = cache.Get("additional_image_2__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 2").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update additional_image_2__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("additional_image_2__productcontenttype", c)
	}

	c = cache.Get("additional_image_3__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 3").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update additional_image_3__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("additional_image_3__productcontenttype", c)
	}

	c = cache.Get("additional_image_4__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 4").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update additional_image_4__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("additional_image_4__productcontenttype", c)
	}

	c = cache.Get("xtra_img_1_small__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_1_small__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_1_small__productcontenttype", c)
	}

	c = cache.Get("xtra_img_1_medium__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_1_medium__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_1_medium__productcontenttype", c)
	}

	c = cache.Get("xtra_img_1_large__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_1_large__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_1_large__productcontenttype", c)
	}

	c = cache.Get("xtra_img_1_detail__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_1_detail__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_1_detail__productcontenttype", c)
	}

	c = cache.Get("xtra_img_2_small__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_2_small__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_2_small__productcontenttype", c)
	}

	c = cache.Get("xtra_img_2_medium__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_2_medium__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_2_medium__productcontenttype", c)
	}

	c = cache.Get("xtra_img_2_large__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_2_large__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_2_large__productcontenttype", c)
	}

	c = cache.Get("xtra_img_2_detail__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_2_detail__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_2_detail__productcontenttype", c)
	}

	c = cache.Get("xtra_img_3_small__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_3_small__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_3_small__productcontenttype", c)
	}

	c = cache.Get("xtra_img_3_medium__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_3_medium__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_3_medium__productcontenttype", c)
	}

	c = cache.Get("xtra_img_3_large__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_3_large__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_3_large__productcontenttype", c)
	}

	c = cache.Get("xtra_img_3_detail__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_3_detail__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_3_detail__productcontenttype", c)
	}

	c = cache.Get("xtra_img_4_small__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_4_small__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_4_small__productcontenttype", c)
	}

	c = cache.Get("xtra_img_4_medium__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_4_medium__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_4_medium__productcontenttype", c)
	}

	c = cache.Get("xtra_img_4_large__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_4_large__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_4_large__productcontenttype", c)
	}

	c = cache.Get("xtra_img_4_detail__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_4_detail__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_4_detail__productcontenttype", c)
	}

	c = cache.Get("xtra_img_l_small__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_l_small__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_l_small__productcontenttype", c)
	}

	c = cache.Get("xtra_img_l_medium__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_l_medium__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_l_medium__productcontenttype", c)
	}

	c = cache.Get("xtra_img_l_large__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_l_large__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_l_large__productcontenttype", c)
	}

	c = cache.Get("xtra_img_l_detail__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image - Additional View Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update xtra_img_l_detail__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("xtra_img_l_detail__productcontenttype", c)
	}

	c = cache.Get("addtocart_label__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Add To Cart Label").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update addtocart_label__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("addtocart_label__productcontenttype", c)
	}

	c = cache.Get("addtocart_image__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Add To Cart Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update addtocart_image__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("addtocart_image__productcontenttype", c)
	}

	c = cache.Get("short_sales_pitch__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Short Sales Pitch").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update short_sales_pitch__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("short_sales_pitch__productcontenttype", c)
	}

	c = cache.Get("installation__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Installation").
		SetParent(cache.Get("digital_download__productcontenttype").(*ent.ProductContentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update installation__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("installation__productcontenttype", c)
	}

	c = cache.Get("specification__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Specification").
		SetParent(cache.Get("digital_download__productcontenttype").(*ent.ProductContentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update specification__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("specification__productcontenttype", c)
	}

	c = cache.Get("warranty__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Warranty").
		SetParent(cache.Get("digital_download__productcontenttype").(*ent.ProductContentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update warranty__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("warranty__productcontenttype", c)
	}

	c = cache.Get("page_title__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Page Title").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update page_title__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("page_title__productcontenttype", c)
	}

	c = cache.Get("meta_keyword__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meta Keyword").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update meta_keyword__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("meta_keyword__productcontenttype", c)
	}

	c = cache.Get("meta_description__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meta Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update meta_description__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("meta_description__productcontenttype", c)
	}

	c = cache.Get("image__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image__productcontenttype", c)
	}

	c = cache.Get("default_image__productcontenttype").(*ent.ProductContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Default Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update default_image__productcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("default_image__productcontenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
