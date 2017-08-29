// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_Comments struct {
	// Authors
	Authors *CT_Authors
	// List of Comments
	CommentList *CT_CommentList
	ExtLst      *CT_ExtensionList
}

func NewCT_Comments() *CT_Comments {
	ret := &CT_Comments{}
	ret.Authors = NewCT_Authors()
	ret.CommentList = NewCT_CommentList()
	return ret
}
func (m *CT_Comments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seauthors := xml.StartElement{Name: xml.Name{Local: "x:authors"}}
	e.EncodeElement(m.Authors, seauthors)
	secommentList := xml.StartElement{Name: xml.Name{Local: "x:commentList"}}
	e.EncodeElement(m.CommentList, secommentList)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Comments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Authors = NewCT_Authors()
	m.CommentList = NewCT_CommentList()
lCT_Comments:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "authors":
				if err := d.DecodeElement(m.Authors, &el); err != nil {
					return err
				}
			case "commentList":
				if err := d.DecodeElement(m.CommentList, &el); err != nil {
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
			break lCT_Comments
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Comments) Validate() error {
	return m.ValidateWithPath("CT_Comments")
}
func (m *CT_Comments) ValidateWithPath(path string) error {
	if err := m.Authors.ValidateWithPath(path + "/Authors"); err != nil {
		return err
	}
	if err := m.CommentList.ValidateWithPath(path + "/CommentList"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}