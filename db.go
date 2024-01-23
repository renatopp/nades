package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/query"
)

type DB struct {
	db *clover.DB
}

func NewDB() *DB {
	db, err := clover.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateCollection("nades")
	db.CreateCollection("variants")

	return &DB{
		db: db,
	}
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) FindNades(mapId string) []Nade {
	nades := []Nade{}

	docs, _ := d.db.FindAll(
		query.NewQuery("nades").
			Where(
				query.Field("map").Eq(mapId),
			),
	)

	for _, doc := range docs {
		nade := Nade{}
		doc.Unmarshal(&nade)
		nades = append(nades, nade)
	}

	return nades
}

func (d *DB) FindVariants(nadeId string) []Variant {
	variants := []Variant{}

	docs, _ := d.db.FindAll(
		query.NewQuery("variants").
			Where(
				query.Field("nade_id").Eq(nadeId),
			),
	)

	for _, doc := range docs {
		variant := Variant{}
		doc.Unmarshal(&variant)
		variants = append(variants, variant)
	}

	return variants
}

func (d *DB) InsertNade(nade Nade) string {
	id, _ := d.db.InsertOne("nades", nade.ToDB())
	return id
}

func (d *DB) InsertVariant(variant Variant) string {
	id, _ := d.db.InsertOne("variants", variant.ToDB())
	return id
}

func (d *DB) UpdateNade(nade Nade) {
	d.db.UpdateById("nades", nade.Id, nade.Update)
}

func (d *DB) UpdateVariant(variant Variant) {
	old, _ := d.db.FindById("variants", variant.Id)
	if old == nil {
		return
	}

	vt := Variant{}
	old.Unmarshal(&vt)

	if vt.Map == "" {
		nodeDoc, _ := d.db.FindById("nades", variant.NadeId)
		m := nodeDoc.Get("map").(string)
		vt.Map = m
		variant.Map = m
	}

	println("Updating", vt.ThrowImage, variant.ThrowImage, vt.Map)
	if vt.ThrowImage != "" && variant.ThrowImage == "" {
		println("Should delete throw image")
		deleteVariantImage(variant.Map, variant.Id, "throw", vt.ThrowImage)
	}
	if vt.LineupImage != "" && variant.LineupImage == "" {
		println("Should delete lineup image")
		deleteVariantImage(variant.Map, variant.Id, "lineup", vt.ThrowImage)
	}
	if vt.PositionImage != "" && variant.PositionImage == "" {
		println("Should delete position image")
		deleteVariantImage(variant.Map, variant.Id, "position", vt.ThrowImage)
	}

	d.db.UpdateById("variants", variant.Id, variant.Update)
}

func (d *DB) DeleteNade(id string) {
	variants := d.FindVariants(id)
	for _, variant := range variants {
		d.DeleteVariant(variant.Id)
	}

	d.db.DeleteById("nades", id)
}

func (d *DB) DeleteVariant(id string) {
	doc, _ := d.db.FindById("variants", id)
	if doc == nil {
		return
	}

	variant := Variant{}
	doc.Unmarshal(&variant)

	doc, _ = d.db.FindById("nades", variant.NadeId)
	if doc != nil {
		nade := Nade{}
		doc.Unmarshal(&nade)

		d := ".screenshots/" + nade.Map + "/" + variant.Id
		os.RemoveAll(d)
	}

	d.db.DeleteById("variants", id)
}

func deleteVariantImage(m, variant, tp, extension string) {
	if !strings.Contains(extension, ".") {
		extension = "." + extension
	}

	d := ".screenshots/" + m + "/" + variant + "/" + tp + extension
	println("Removing", d)
	go func() {
		for i := 0; i < 10; i++ {
			err := os.Remove(d)
			if err != nil {
				println(err.Error())
				time.Sleep(time.Duration(i+1) * time.Second)
				continue
			}
			return
		}
	}()
}
