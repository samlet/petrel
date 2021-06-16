package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductContentType(ctx context.Context) error {
	log.Println("ProductContentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductContentType

	c, err = client.ProductContentType.Create().SetStringRef("online_access__productcontenttype").
		SetHasTable("No").
		SetDescription("Online Access").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create online_access__productcontenttype: %v", err)
		return err
	}
	cache.Put("online_access__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("digital_download__productcontenttype").
		SetHasTable("No").
		SetDescription("Digital Download").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create digital_download__productcontenttype: %v", err)
		return err
	}
	cache.Put("digital_download__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("fulfillment_email__productcontenttype").
		SetHasTable("No").
		SetDescription("Fulfillment Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fulfillment_email__productcontenttype: %v", err)
		return err
	}
	cache.Put("fulfillment_email__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("fulfillment_extasync__productcontenttype").
		SetHasTable("No").
		SetDescription("Fulfillment External (Async)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fulfillment_extasync__productcontenttype: %v", err)
		return err
	}
	cache.Put("fulfillment_extasync__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("fulfillment_extsync__productcontenttype").
		SetHasTable("No").
		SetDescription("Fulfillment External (Sync)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fulfillment_extsync__productcontenttype: %v", err)
		return err
	}
	cache.Put("fulfillment_extsync__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("product_name__productcontenttype").
		SetHasTable("No").
		SetDescription("Product Name").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_name__productcontenttype: %v", err)
		return err
	}
	cache.Put("product_name__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("description__productcontenttype").
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create description__productcontenttype: %v", err)
		return err
	}
	cache.Put("description__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("long_description__productcontenttype").
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create long_description__productcontenttype: %v", err)
		return err
	}
	cache.Put("long_description__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("alternative_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Alternative URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create alternative_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("alternative_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("price_detail_text__productcontenttype").
		SetHasTable("No").
		SetDescription("Price Detail Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_detail_text__productcontenttype: %v", err)
		return err
	}
	cache.Put("price_detail_text__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("ingredients__productcontenttype").
		SetHasTable("No").
		SetDescription("Ingredients").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ingredients__productcontenttype: %v", err)
		return err
	}
	cache.Put("ingredients__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("unique_ingredients__productcontenttype").
		SetHasTable("No").
		SetDescription("Unique Ingredients").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create unique_ingredients__productcontenttype: %v", err)
		return err
	}
	cache.Put("unique_ingredients__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("specialinstructions__productcontenttype").
		SetHasTable("No").
		SetDescription("Special Instructions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create specialinstructions__productcontenttype: %v", err)
		return err
	}
	cache.Put("specialinstructions__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("warnings__productcontenttype").
		SetHasTable("No").
		SetDescription("Warnings").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create warnings__productcontenttype: %v", err)
		return err
	}
	cache.Put("warnings__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("directions__productcontenttype").
		SetHasTable("No").
		SetDescription("Directions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create directions__productcontenttype: %v", err)
		return err
	}
	cache.Put("directions__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("terms_and_conds__productcontenttype").
		SetHasTable("No").
		SetDescription("Terms and Conditions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create terms_and_conds__productcontenttype: %v", err)
		return err
	}
	cache.Put("terms_and_conds__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("delivery_info__productcontenttype").
		SetHasTable("No").
		SetDescription("Delivery Info").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create delivery_info__productcontenttype: %v", err)
		return err
	}
	cache.Put("delivery_info__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("small_image_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create small_image_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("small_image_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("medium_image_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create medium_image_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("medium_image_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("large_image_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create large_image_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("large_image_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("detail_image_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create detail_image_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("detail_image_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("original_image_url__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Original").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create original_image_url__productcontenttype: %v", err)
		return err
	}
	cache.Put("original_image_url__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("small_image_alt__productcontenttype").
		SetHasTable("No").
		SetDescription("Image Alt Text - Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create small_image_alt__productcontenttype: %v", err)
		return err
	}
	cache.Put("small_image_alt__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("medium_image_alt__productcontenttype").
		SetHasTable("No").
		SetDescription("Image Alt Text - Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create medium_image_alt__productcontenttype: %v", err)
		return err
	}
	cache.Put("medium_image_alt__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("large_image_alt__productcontenttype").
		SetHasTable("No").
		SetDescription("Image Alt Text - Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create large_image_alt__productcontenttype: %v", err)
		return err
	}
	cache.Put("large_image_alt__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("detail_image_alt__productcontenttype").
		SetHasTable("No").
		SetDescription("Image Alt Text - Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create detail_image_alt__productcontenttype: %v", err)
		return err
	}
	cache.Put("detail_image_alt__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("additional_image_1__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create additional_image_1__productcontenttype: %v", err)
		return err
	}
	cache.Put("additional_image_1__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("additional_image_2__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 2").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create additional_image_2__productcontenttype: %v", err)
		return err
	}
	cache.Put("additional_image_2__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("additional_image_3__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 3").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create additional_image_3__productcontenttype: %v", err)
		return err
	}
	cache.Put("additional_image_3__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("additional_image_4__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 4").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create additional_image_4__productcontenttype: %v", err)
		return err
	}
	cache.Put("additional_image_4__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_1_small__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_1_small__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_1_small__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_1_medium__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_1_medium__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_1_medium__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_1_large__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_1_large__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_1_large__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_1_detail__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 1 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_1_detail__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_1_detail__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_2_small__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_2_small__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_2_small__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_2_medium__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_2_medium__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_2_medium__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_2_large__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_2_large__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_2_large__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_2_detail__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 2 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_2_detail__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_2_detail__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_3_small__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_3_small__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_3_small__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_3_medium__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_3_medium__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_3_medium__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_3_large__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_3_large__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_3_large__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_3_detail__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 3 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_3_detail__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_3_detail__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_4_small__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_4_small__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_4_small__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_4_medium__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_4_medium__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_4_medium__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_4_large__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_4_large__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_4_large__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_4_detail__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View 4 Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_4_detail__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_4_detail__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_l_small__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View Small").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_l_small__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_l_small__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_l_medium__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View Medium").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_l_medium__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_l_medium__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_l_large__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View Large").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_l_large__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_l_large__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("xtra_img_l_detail__productcontenttype").
		SetHasTable("No").
		SetDescription("Image - Additional View Detail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create xtra_img_l_detail__productcontenttype: %v", err)
		return err
	}
	cache.Put("xtra_img_l_detail__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("addtocart_label__productcontenttype").
		SetHasTable("No").
		SetDescription("Add To Cart Label").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create addtocart_label__productcontenttype: %v", err)
		return err
	}
	cache.Put("addtocart_label__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("addtocart_image__productcontenttype").
		SetHasTable("No").
		SetDescription("Add To Cart Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create addtocart_image__productcontenttype: %v", err)
		return err
	}
	cache.Put("addtocart_image__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("short_sales_pitch__productcontenttype").
		SetHasTable("No").
		SetDescription("Short Sales Pitch").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create short_sales_pitch__productcontenttype: %v", err)
		return err
	}
	cache.Put("short_sales_pitch__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("installation__productcontenttype").
		SetHasTable("No").
		SetDescription("Installation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create installation__productcontenttype: %v", err)
		return err
	}
	cache.Put("installation__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("specification__productcontenttype").
		SetHasTable("No").
		SetDescription("Specification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create specification__productcontenttype: %v", err)
		return err
	}
	cache.Put("specification__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("warranty__productcontenttype").
		SetHasTable("No").
		SetDescription("Warranty").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create warranty__productcontenttype: %v", err)
		return err
	}
	cache.Put("warranty__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("page_title__productcontenttype").
		SetHasTable("No").
		SetDescription("Page Title").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create page_title__productcontenttype: %v", err)
		return err
	}
	cache.Put("page_title__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("meta_keyword__productcontenttype").
		SetHasTable("No").
		SetDescription("Meta Keyword").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create meta_keyword__productcontenttype: %v", err)
		return err
	}
	cache.Put("meta_keyword__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("meta_description__productcontenttype").
		SetHasTable("No").
		SetDescription("Meta Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create meta_description__productcontenttype: %v", err)
		return err
	}
	cache.Put("meta_description__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("image__productcontenttype").
		SetHasTable("No").
		SetDescription("Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image__productcontenttype: %v", err)
		return err
	}
	cache.Put("image__productcontenttype", c)

	c, err = client.ProductContentType.Create().SetStringRef("default_image__productcontenttype").
		SetHasTable("No").
		SetDescription("Default Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create default_image__productcontenttype: %v", err)
		return err
	}
	cache.Put("default_image__productcontenttype", c)

	return nil
}
