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
	"strconv"
)

type CT_SingleXmlCell struct {
	// Table Id
	IdAttr uint32
	// Reference
	RAttr string
	// Connection ID
	ConnectionIdAttr uint32
	// Cell Properties
	XmlCellPr *CT_XmlCellPr
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_SingleXmlCell() *CT_SingleXmlCell {
	ret := &CT_SingleXmlCell{}
	ret.XmlCellPr = NewCT_XmlCellPr()
	return ret
}
func (m *CT_SingleXmlCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectionId"},
		Value: fmt.Sprintf("%v", m.ConnectionIdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	sexmlCellPr := xml.StartElement{Name: xml.Name{Local: "x:xmlCellPr"}}
	e.EncodeElement(m.XmlCellPr, sexmlCellPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SingleXmlCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.XmlCellPr = NewCT_XmlCellPr()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
		}
		if attr.Name.Local == "connectionId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ConnectionIdAttr = uint32(parsed)
		}
	}
lCT_SingleXmlCell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "xmlCellPr":
				if err := d.DecodeElement(m.XmlCellPr, &el); err != nil {
					return err
				}
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
			break lCT_SingleXmlCell
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SingleXmlCell) Validate() error {
	return m.ValidateWithPath("CT_SingleXmlCell")
}
func (m *CT_SingleXmlCell) ValidateWithPath(path string) error {
	if err := m.XmlCellPr.ValidateWithPath(path + "/XmlCellPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}