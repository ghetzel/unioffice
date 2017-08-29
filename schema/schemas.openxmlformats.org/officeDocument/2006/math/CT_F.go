// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_F struct {
	FPr *CT_FPr
	Num *CT_OMathArg
	Den *CT_OMathArg
}

func NewCT_F() *CT_F {
	ret := &CT_F{}
	ret.Num = NewCT_OMathArg()
	ret.Den = NewCT_OMathArg()
	return ret
}
func (m *CT_F) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.FPr != nil {
		sefPr := xml.StartElement{Name: xml.Name{Local: "m:fPr"}}
		e.EncodeElement(m.FPr, sefPr)
	}
	senum := xml.StartElement{Name: xml.Name{Local: "m:num"}}
	e.EncodeElement(m.Num, senum)
	seden := xml.StartElement{Name: xml.Name{Local: "m:den"}}
	e.EncodeElement(m.Den, seden)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_F) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Num = NewCT_OMathArg()
	m.Den = NewCT_OMathArg()
lCT_F:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fPr":
				m.FPr = NewCT_FPr()
				if err := d.DecodeElement(m.FPr, &el); err != nil {
					return err
				}
			case "num":
				if err := d.DecodeElement(m.Num, &el); err != nil {
					return err
				}
			case "den":
				if err := d.DecodeElement(m.Den, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_F
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_F) Validate() error {
	return m.ValidateWithPath("CT_F")
}
func (m *CT_F) ValidateWithPath(path string) error {
	if m.FPr != nil {
		if err := m.FPr.ValidateWithPath(path + "/FPr"); err != nil {
			return err
		}
	}
	if err := m.Num.ValidateWithPath(path + "/Num"); err != nil {
		return err
	}
	if err := m.Den.ValidateWithPath(path + "/Den"); err != nil {
		return err
	}
	return nil
}