package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductMeterType(ctx context.Context) error {
	log.Println("ProductMeterType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductMeterType
	failures := 0

	c = cache.Get("speedometer__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Speedometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update speedometer__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("speedometer__productmetertype", c)
	}

	c = cache.Get("tachometer__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Tachometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tachometer__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tachometer__productmetertype", c)
	}

	c = cache.Get("distance__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Distance Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update distance__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("distance__productmetertype", c)
	}

	c = cache.Get("motor_time__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Motor Time Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update motor_time__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("motor_time__productmetertype", c)
	}

	c = cache.Get("use_count__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Use Count Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update use_count__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("use_count__productmetertype", c)
	}

	c = cache.Get("copy_count__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Copy Count Meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update copy_count__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("copy_count__productmetertype", c)
	}

	c = cache.Get("trip_meter__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Trip meter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update trip_meter__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("trip_meter__productmetertype", c)
	}

	c = cache.Get("tachograph__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Tachograph").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tachograph__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tachograph__productmetertype", c)
	}

	c = cache.Get("taximeter__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Taximeter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update taximeter__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("taximeter__productmetertype", c)
	}

	c = cache.Get("event_data_recorder__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Event Data Recorder").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update event_data_recorder__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("event_data_recorder__productmetertype", c)
	}

	c = cache.Get("pedometer__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Pedometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pedometer__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pedometer__productmetertype", c)
	}

	c = cache.Get("odometer__productmetertype").(*ent.ProductMeterType)
	c, err = c.Update().
		SetDescription("Odometer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update odometer__productmetertype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("odometer__productmetertype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
