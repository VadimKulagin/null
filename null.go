// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import "encoding/xml"

const jsonNull = "null"

var xsiNilAttr = xml.Attr{
	Name:  xml.Name{Local: "xsi:nil"},
	Value: "true",
}

func isXsiNilAttr(attr xml.Attr) bool {
	return attr.Name.Space == "xsi" && attr.Name.Local == "nil" && attr.Value == "true"
}
