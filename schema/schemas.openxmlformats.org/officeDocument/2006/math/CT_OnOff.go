// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_OnOff struct {
	ValAttr *sharedTypes.ST_OnOff
}

func NewCT_OnOff() *CT_OnOff {
	ret := &CT_OnOff{}
	return ret
}
func (m *CT_OnOff) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "m:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OnOff) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OnOff: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_OnOff) Validate() error {
	return m.ValidateWithPath("CT_OnOff")
}
func (m *CT_OnOff) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
			return err
		}
	}
	return nil
}