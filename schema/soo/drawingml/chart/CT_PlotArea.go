// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/soo/drawingml"
)

type CT_PlotArea struct {
	Layout  *CT_Layout
	Choice  []*CT_PlotAreaChoice
	CChoice *CT_PlotAreaChoice1
	DTable  *CT_DTable
	SpPr    *drawingml.CT_ShapeProperties
	ExtLst  *CT_ExtensionList
}

func NewCT_PlotArea() *CT_PlotArea {
	ret := &CT_PlotArea{}
	return ret
}

func (m *CT_PlotArea) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Layout != nil {
		selayout := xml.StartElement{Name: xml.Name{Local: "c:layout"}}
		e.EncodeElement(m.Layout, selayout)
	}
	for _, c := range m.Choice {
		c.MarshalXML(e, xml.StartElement{})
	}
	if m.CChoice != nil {
		m.CChoice.MarshalXML(e, xml.StartElement{})
	}
	if m.DTable != nil {
		sedTable := xml.StartElement{Name: xml.Name{Local: "c:dTable"}}
		e.EncodeElement(m.DTable, sedTable)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PlotArea) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PlotArea:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "layout"}:
				m.Layout = NewCT_Layout()
				if err := d.DecodeElement(m.Layout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "areaChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.AreaChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "area3DChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.Area3DChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "lineChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.LineChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "line3DChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.Line3DChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "stockChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.StockChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "radarChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.RadarChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "scatterChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.ScatterChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pieChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.PieChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pie3DChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.Pie3DChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "doughnutChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.DoughnutChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "barChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.BarChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bar3DChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.Bar3DChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ofPieChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.OfPieChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "surfaceChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.SurfaceChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "surface3DChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.Surface3DChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bubbleChart"}:
				tmp := NewCT_PlotAreaChoice()
				if err := d.DecodeElement(&tmp.BubbleChart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "valAx"}:
				m.CChoice = NewCT_PlotAreaChoice1()
				if err := d.DecodeElement(&m.CChoice.ValAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "catAx"}:
				m.CChoice = NewCT_PlotAreaChoice1()
				if err := d.DecodeElement(&m.CChoice.CatAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dateAx"}:
				m.CChoice = NewCT_PlotAreaChoice1()
				if err := d.DecodeElement(&m.CChoice.DateAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "serAx"}:
				m.CChoice = NewCT_PlotAreaChoice1()
				if err := d.DecodeElement(&m.CChoice.SerAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dTable"}:
				m.DTable = NewCT_DTable()
				if err := d.DecodeElement(m.DTable, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PlotArea %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PlotArea
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PlotArea and its children
func (m *CT_PlotArea) Validate() error {
	return m.ValidateWithPath("CT_PlotArea")
}

// ValidateWithPath validates the CT_PlotArea and its children, prefixing error messages with path
func (m *CT_PlotArea) ValidateWithPath(path string) error {
	if m.Layout != nil {
		if err := m.Layout.ValidateWithPath(path + "/Layout"); err != nil {
			return err
		}
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CChoice != nil {
		if err := m.CChoice.ValidateWithPath(path + "/CChoice"); err != nil {
			return err
		}
	}
	if m.DTable != nil {
		if err := m.DTable.ValidateWithPath(path + "/DTable"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}