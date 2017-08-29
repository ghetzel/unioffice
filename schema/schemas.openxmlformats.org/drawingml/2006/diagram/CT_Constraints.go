// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Constraints struct {
	Constr []*CT_Constraint
}

func NewCT_Constraints() *CT_Constraints {
	ret := &CT_Constraints{}
	return ret
}
func (m *CT_Constraints) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Constr != nil {
		seconstr := xml.StartElement{Name: xml.Name{Local: "constr"}}
		e.EncodeElement(m.Constr, seconstr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Constraints) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Constraints:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "constr":
				tmp := NewCT_Constraint()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Constr = append(m.Constr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Constraints
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Constraints) Validate() error {
	return m.ValidateWithPath("CT_Constraints")
}
func (m *CT_Constraints) ValidateWithPath(path string) error {
	for i, v := range m.Constr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Constr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}