package main

import (
	"fmt"

	"github.com/ostafen/clover/v2/document"
)

type NadeType string

const SMOKE NadeType = "smoke"
const FLASH NadeType = "flash"
const MOLLY NadeType = "molly"
const HE NadeType = "he"

type Nade struct {
	Id   string   `clover:"_id" json:"id"`
	Map  string   `clover:"map" json:"map"`
	Type NadeType `clover:"type" json:"type"`
	X    int      `clover:"x" json:"x"`
	Y    int      `clover:"y" json:"y"`
}

func (n *Nade) ToDB() *document.Document {
	doc := document.NewDocument()
	doc.Set("map", n.Map)
	doc.Set("type", n.Type)
	doc.Set("x", n.X)
	doc.Set("y", n.Y)
	return doc
}

func (n *Nade) Update(doc *document.Document) *document.Document {
	doc.Set("x", n.X)
	doc.Set("y", n.Y)
	return doc
}

func (n *Nade) ToString() string {
	return fmt.Sprintf("[%s] (%d, %d)", n.Type, n.X, n.Y)
}

type Variant struct {
	Id            string `clover:"_id" json:"id"`
	NadeId        string `clover:"nade_id" json:"nadeId"`
	Name          string `clover:"name" json:"name"`
	Description   string `clover:"description" json:"description"`
	X             int    `clover:"x" json:"x"`
	Y             int    `clover:"y" json:"y"`
	ThrowImage    bool   `clover:"throw_image" json:"throwImage"`
	LineupImage   bool   `clover:"lineup_image" json:"lineupImage"`
	PositionImage bool   `clover:"position_image" json:"positionImage"`
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

func (v *Variant) Update(doc *document.Document) *document.Document {
	doc.Set("name", v.Name)
	doc.Set("description", v.Description)
	doc.Set("x", v.X)
	doc.Set("y", v.Y)
	doc.Set("throw_image", v.ThrowImage)
	doc.Set("lineup_image", v.LineupImage)
	doc.Set("position_image", v.PositionImage)
	return doc
}

func (v *Variant) ToString() string {
	return fmt.Sprintf("[%s] %s (%d, %d)", v.NadeId, v.Name, v.X, v.Y)
}
