// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_XYAdjustHandle struct {
	GdRefXAttr *string
	MinXAttr   *ST_AdjCoordinate
	MaxXAttr   *ST_AdjCoordinate
	GdRefYAttr *string
	MinYAttr   *ST_AdjCoordinate
	MaxYAttr   *ST_AdjCoordinate
	Pos        *CT_AdjPoint2D
}

func NewCT_XYAdjustHandle() *CT_XYAdjustHandle {
	ret := &CT_XYAdjustHandle{}
	ret.Pos = NewCT_AdjPoint2D()
	return ret
}
func (m *CT_XYAdjustHandle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.GdRefXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gdRefX"},
			Value: fmt.Sprintf("%v", *m.GdRefXAttr)})
	}
	if m.MinXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minX"},
			Value: fmt.Sprintf("%v", *m.MinXAttr)})
	}
	if m.MaxXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxX"},
			Value: fmt.Sprintf("%v", *m.MaxXAttr)})
	}
	if m.GdRefYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gdRefY"},
			Value: fmt.Sprintf("%v", *m.GdRefYAttr)})
	}
	if m.MinYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minY"},
			Value: fmt.Sprintf("%v", *m.MinYAttr)})
	}
	if m.MaxYAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxY"},
			Value: fmt.Sprintf("%v", *m.MaxYAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sepos := xml.StartElement{Name: xml.Name{Local: "a:pos"}}
	e.EncodeElement(m.Pos, sepos)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_XYAdjustHandle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = NewCT_AdjPoint2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "gdRefX" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GdRefXAttr = &parsed
		}
		if attr.Name.Local == "minX" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MinXAttr = &parsed
		}
		if attr.Name.Local == "maxX" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MaxXAttr = &parsed
		}
		if attr.Name.Local == "gdRefY" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GdRefYAttr = &parsed
		}
		if attr.Name.Local == "minY" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MinYAttr = &parsed
		}
		if attr.Name.Local == "maxY" {
			parsed, err := ParseUnionST_AdjCoordinate(attr.Value)
			if err != nil {
				return err
			}
			m.MaxYAttr = &parsed
		}
	}
lCT_XYAdjustHandle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "pos":
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_XYAdjustHandle
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_XYAdjustHandle) Validate() error {
	return m.ValidateWithPath("CT_XYAdjustHandle")
}
func (m *CT_XYAdjustHandle) ValidateWithPath(path string) error {
	if m.MinXAttr != nil {
		if err := m.MinXAttr.ValidateWithPath(path + "/MinXAttr"); err != nil {
			return err
		}
	}
	if m.MaxXAttr != nil {
		if err := m.MaxXAttr.ValidateWithPath(path + "/MaxXAttr"); err != nil {
			return err
		}
	}
	if m.MinYAttr != nil {
		if err := m.MinYAttr.ValidateWithPath(path + "/MinYAttr"); err != nil {
			return err
		}
	}
	if m.MaxYAttr != nil {
		if err := m.MaxYAttr.ValidateWithPath(path + "/MaxYAttr"); err != nil {
			return err
		}
	}
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	return nil
}