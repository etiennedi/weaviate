/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

package gremlin

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/creativesoftwarefdn/weaviate/models"
)

// keyToGremlin translates a thing struct into a vertex string
func (f *Gremlin) keyToGremlin(objToHandle *models.Key, token string, UUID string, addOrUpdate string) (string, string) {

	// create vertex and edge buffers
	var vertex bytes.Buffer
	var edge bytes.Buffer

	if addOrUpdate == "add" {
		// set general fields
		vertex.WriteString(`g.addV("key").property("uuid", "` + UUID + `").property("type", "key")`)
	} else if addOrUpdate == "update" {
		vertex.WriteString(`g.V().hasLabel("key")`)
		vertex.WriteString(`.has("uuid", "` + UUID + `").has("type", "key")`)
	} else {
		f.messaging.ErrorMessage("Trying to add or update a key without `addOrUpdate` set")
		return "", ""
	}

	// set boolean properties
	isRoot := objToHandle.Parent == nil
	vertex.WriteString(`.property("isRoot", ` + strconv.FormatBool(isRoot) + `)`)
	vertex.WriteString(`.property("delete", ` + strconv.FormatBool(objToHandle.Delete) + `)`)
	vertex.WriteString(`.property("execute", ` + strconv.FormatBool(objToHandle.Execute) + `)`)
	vertex.WriteString(`.property("read", ` + strconv.FormatBool(objToHandle.Read) + `)`)
	vertex.WriteString(`.property("write", ` + strconv.FormatBool(objToHandle.Write) + `)`)

	// set string properties
	vertex.WriteString(`.property("email", "` + objToHandle.Email + `")`)

	// set array properties
	vertex.WriteString(`.property("IPOrigin", "` + strings.Join(objToHandle.IPOrigin, ";") + `")`)

	// set integers
	vertex.WriteString(`.property("keyExpiresUnix", ` + strconv.FormatInt(objToHandle.KeyExpiresUnix, 10) + `)`)

	// add the secured and encrypted token
	vertex.WriteString(`.property("__token", "` + base64.StdEncoding.EncodeToString([]byte(token)) + `")`)

	// if isRoot is false, an edge needs to be added
	if isRoot == false {
		edge.WriteString(`g.addE("parent").from(g.V().hasLabel("key").has("uuid", "` + UUID + `")).to(g.V().hasLabel("key").has("uuid", "` + objToHandle.Parent.NrDollarCref.String() + `")).property("\$cref", "` + objToHandle.Parent.NrDollarCref.String() + `").property("type", "` + objToHandle.Parent.Type + `").property("locationUrl", "` + *objToHandle.Parent.LocationURL + `")`)
	}

	return vertex.String(), edge.String()
}
