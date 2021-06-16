package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductMeterType(ctx context.Context) error {
	log.Println("ProductMeterType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductMeterType

	c, err = client.ProductMeterType.Create().SetStringRef("speedometer__productmetertype").
		SetDescription("Speedometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create speedometer__productmetertype: %v", err)
		return err
	}
	cache.Put("speedometer__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("tachometer__productmetertype").
		SetDescription("Tachometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tachometer__productmetertype: %v", err)
		return err
	}
	cache.Put("tachometer__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("distance__productmetertype").
		SetDescription("Distance Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distance__productmetertype: %v", err)
		return err
	}
	cache.Put("distance__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("motor_time__productmetertype").
		SetDescription("Motor Time Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create motor_time__productmetertype: %v", err)
		return err
	}
	cache.Put("motor_time__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("use_count__productmetertype").
		SetDescription("Use Count Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create use_count__productmetertype: %v", err)
		return err
	}
	cache.Put("use_count__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("copy_count__productmetertype").
		SetDescription("Copy Count Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create copy_count__productmetertype: %v", err)
		return err
	}
	cache.Put("copy_count__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("trip_meter__productmetertype").
		SetDescription("Trip meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create trip_meter__productmetertype: %v", err)
		return err
	}
	cache.Put("trip_meter__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("tachograph__productmetertype").
		SetDescription("Tachograph").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tachograph__productmetertype: %v", err)
		return err
	}
	cache.Put("tachograph__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("taximeter__productmetertype").
		SetDescription("Taximeter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create taximeter__productmetertype: %v", err)
		return err
	}
	cache.Put("taximeter__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("event_data_recorder__productmetertype").
		SetDescription("Event Data Recorder").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create event_data_recorder__productmetertype: %v", err)
		return err
	}
	cache.Put("event_data_recorder__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("pedometer__productmetertype").
		SetDescription("Pedometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pedometer__productmetertype: %v", err)
		return err
	}
	cache.Put("pedometer__productmetertype", c)

	c, err = client.ProductMeterType.Create().SetStringRef("odometer__productmetertype").
		SetDescription("Odometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create odometer__productmetertype: %v", err)
		return err
	}
	cache.Put("odometer__productmetertype", c)

	return nil
}
