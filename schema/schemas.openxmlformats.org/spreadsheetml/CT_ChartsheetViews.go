// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_ChartsheetViews struct {
	// Chart Sheet View
	SheetView []*CT_ChartsheetView
	ExtLst    *CT_ExtensionList
}

func NewCT_ChartsheetViews() *CT_ChartsheetViews {
	ret := &CT_ChartsheetViews{}
	return ret
}
func (m *CT_ChartsheetViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sesheetView := xml.StartElement{Name: xml.Name{Local: "x:sheetView"}}
	e.EncodeElement(m.SheetView, sesheetView)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ChartsheetViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ChartsheetViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetView":
				tmp := NewCT_ChartsheetView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SheetView = append(m.SheetView, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetViews
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ChartsheetViews) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetViews")
}
func (m *CT_ChartsheetViews) ValidateWithPath(path string) error {
	for i, v := range m.SheetView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SheetView[%d]", path, i)); err != nil {
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