// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TLTimeAnimateValueList struct {
	// Time Animate Value
	Tav []*CT_TLTimeAnimateValue
}

func NewCT_TLTimeAnimateValueList() *CT_TLTimeAnimateValueList {
	ret := &CT_TLTimeAnimateValueList{}
	return ret
}
func (m *CT_TLTimeAnimateValueList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Tav != nil {
		setav := xml.StartElement{Name: xml.Name{Local: "p:tav"}}
		e.EncodeElement(m.Tav, setav)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLTimeAnimateValueList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTimeAnimateValueList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tav":
				tmp := NewCT_TLTimeAnimateValue()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tav = append(m.Tav, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeAnimateValueList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLTimeAnimateValueList) Validate() error {
	return m.ValidateWithPath("CT_TLTimeAnimateValueList")
}
func (m *CT_TLTimeAnimateValueList) ValidateWithPath(path string) error {
	for i, v := range m.Tav {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tav[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}