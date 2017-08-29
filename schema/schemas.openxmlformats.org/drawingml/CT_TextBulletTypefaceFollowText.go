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
)

type CT_TextBulletTypefaceFollowText struct {
}

func NewCT_TextBulletTypefaceFollowText() *CT_TextBulletTypefaceFollowText {
	ret := &CT_TextBulletTypefaceFollowText{}
	return ret
}
func (m *CT_TextBulletTypefaceFollowText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextBulletTypefaceFollowText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextBulletTypefaceFollowText: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TextBulletTypefaceFollowText) Validate() error {
	return m.ValidateWithPath("CT_TextBulletTypefaceFollowText")
}
func (m *CT_TextBulletTypefaceFollowText) ValidateWithPath(path string) error {
	return nil
}