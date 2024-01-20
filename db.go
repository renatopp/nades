package main

import (
	"log"

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
	d.db.UpdateById("variants", variant.Id, variant.Update)
}

func (d *DB) DeleteNade(id string) {
	d.db.DeleteById("nades", id)
}

func (d *DB) DeleteVariant(id string) {
	d.db.DeleteById("variants", id)
}
