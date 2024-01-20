package main

import (
	"fmt"
	"log"

	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

type Nade struct {
	MapId string `clover:"map_id"`
	Name  string `clover:"name"`
	X     int    `clover:"x"`
	Y     int    `clover:"y"`
}

func (n *Nade) ToDB() *document.Document {
	doc := document.NewDocument()
	doc.Set("map_id", n.MapId)
	doc.Set("name", n.Name)
	doc.Set("x", n.X)
	doc.Set("y", n.Y)
	return doc
}

func (n *Nade) ToString() string {
	return fmt.Sprintf("[%s] %s (%d, %d)", n.MapId, n.Name, n.X, n.Y)
}

type Variant struct {
	NadeId      string `clover:"nade_id"`
	Name        string `clover:"name"`
	Description string `clover:"description"`
	X           int    `clover:"x"`
	Y           int    `clover:"y"`
}

func (v *Variant) ToDB() *document.Document {
	doc := document.NewDocument()
	doc.Set("nade_id", v.NadeId)
	doc.Set("name", v.Name)
	doc.Set("description", v.Description)
	doc.Set("x", v.X)
	doc.Set("y", v.Y)
	return doc
}

func (v *Variant) ToString() string {
	return fmt.Sprintf("[%s] %s (%d, %d)", v.NadeId, v.Name, v.X, v.Y)
}

var db *clover.DB

func GetNades(mapId string) []Nade {
	nades := []Nade{}

	docs, _ := db.FindAll(
		query.
			NewQuery("nades").
			Where(
				query.Field("map_id").Eq(mapId),
			),
	)
	for _, doc := range docs {
		nade := Nade{}
		doc.Unmarshal(&nade)
		nades = append(nades, nade)
	}

	return nades
}

func main() {
	_db, err := clover.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer _db.Close()
	db = _db

	db.CreateCollection("nades")
	db.CreateCollection("variants")

	nade := Nade{
		MapId: "de_dust2",
		Name:  "Smoke Mid",
		X:     100,
		Y:     200,
	}
	db.InsertOne("nades", nade.ToDB())

	nade = Nade{
		MapId: "de_dust2",
		Name:  "Smoke Long",
		X:     500,
		Y:     100,
	}
	db.InsertOne("nades", nade.ToDB())

	nade = Nade{
		MapId: "de_mirage",
		Name:  "Smoke Long",
		X:     500,
		Y:     100,
	}
	db.InsertOne("nades", nade.ToDB())

	nades := GetNades("de_dust2")
	for _, nade := range nades {
		log.Println(nade.ToString())
	}
}
