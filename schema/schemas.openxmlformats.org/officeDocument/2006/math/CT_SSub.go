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

type CT_SSub struct {
	SSubPr *CT_SSubPr
	E      *CT_OMathArg
	Sub    *CT_OMathArg
}

func NewCT_SSub() *CT_SSub {
	ret := &CT_SSub{}
	ret.E = NewCT_OMathArg()
	ret.Sub = NewCT_OMathArg()
	return ret
}
func (m *CT_SSub) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.SSubPr != nil {
		sesSubPr := xml.StartElement{Name: xml.Name{Local: "m:sSubPr"}}
		e.EncodeElement(m.SSubPr, sesSubPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	sesub := xml.StartElement{Name: xml.Name{Local: "m:sub"}}
	e.EncodeElement(m.Sub, sesub)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Sub = NewCT_OMathArg()
lCT_SSub:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sSubPr":
				m.SSubPr = NewCT_SSubPr()
				if err := d.DecodeElement(m.SSubPr, &el); err != nil {
					return err
				}
			case "e":
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case "sub":
				if err := d.DecodeElement(m.Sub, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSub
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SSub) Validate() error {
	return m.ValidateWithPath("CT_SSub")
}
func (m *CT_SSub) ValidateWithPath(path string) error {
	if m.SSubPr != nil {
		if err := m.SSubPr.ValidateWithPath(path + "/SSubPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Sub.ValidateWithPath(path + "/Sub"); err != nil {
		return err
	}
	return nil
}