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

package janusgraph

import (
	"bytes"
	"encoding/json"
	errors_ "errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/creativesoftwarefdn/weaviate/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/models"
)

// processes a single thing
func (f *Janusgraph) processSingleThing(result interface{}, thingNo int, thingResponse *models.ThingGetResponse) error {

	// Create the schema Map, this map will contain all the results
	responseSchema := make(map[string]interface{})

	// set meta values (String)
	thingResponse.ThingID = strfmt.UUID(f.getSinglePropertyValue(result, "uuid", thingNo).(string))
	thingResponse.AtClass = f.getSinglePropertyValue(result, "atClass", thingNo).(string)
	thingResponse.AtContext = f.getSinglePropertyValue(result, "context", thingNo).(string)

	// set meta values (int64)
	thingResponse.CreationTimeUnix, _ = strconv.ParseInt(f.getSinglePropertyValue(result, "creationTimeUnix", thingNo).(string), 10, 64)
	thingResponse.LastUpdateTimeUnix, _ = strconv.ParseInt(f.getSinglePropertyValue(result, "lastUpdateTimeUnix", thingNo).(string), 10, 64)

	// Loop over the Janusgraph schema, results
	for key, value := range result.([]interface{})[0].([]interface{})[thingNo].(map[string]interface{}) {
		// if a complete schema is submitted
		if key == "properties" {
			for propKey, propValue := range value.(map[string]interface{}) {
				// check if the key starts with schema__ prefix
				if strings.HasPrefix(propKey, "schema__") {
					// Grab the value and valueType of the properties
					for _, propValueSingle := range propValue.([]interface{}) {
						// loop over the id's and the values, add the interface{} type to the response schema
						for propValueSingleKey, propValueSingleValue := range propValueSingle.(map[string]interface{}) {
							if propValueSingleKey == "value" {
								responseSchema[propKey[8:]] = propValueSingleValue
							}
						}
					}
				}
			}
		}

		// find the edges (if any)
		result, err := f.client.Execute(
			`g.V().hasLabel("thing").has("uuid", "`+string(thingResponse.ThingID)+`").outE()`,
			map[string]string{},
			map[string]string{},
		)

		// in case of error, return the error
		if err != nil {
			return err
		}

		// define Edges
		var edges Edges

		// edges to JSON
		edgesBytes, err := json.Marshal(result)

		// in case of error, return the error
		if err != nil {
			return err
		}

		// merge edges
		err = json.Unmarshal(edgesBytes, &edges)

		// in case of error, return the error
		if err != nil {
			return err
		}

		// add the properties to the edges. Note that the NrDollarCref is _not_ the Janusgraph/Janusgraph ID but the Weaviate UUID
		for _, edgeValue := range edges[0] {
			if len(edgeValue.Label) >= 8 { // should be larger than 8
				if edgeValue.Label[0:8] == "schema__" { // only handle schema edges
					responseSchema[edgeValue.Label[8:]] = models.SingleRef{
						NrDollarCref: edgeValue.Properties.NrDollarCref,
						Type:         edgeValue.Properties.Type,
						LocationURL:  edgeValue.Properties.LocationURL,
					}
				}
			} else if edgeValue.Label == "key" { // smaller then 8 and = "key"
				// get the related key and return
				keyUUID, err := f.keyUUIDFromEdge(edgeValue)

				// in case of error, return the error
				if err != nil {
					return err
				}

				thingResponse.Key = f.createCrefObject(keyUUID, f.serverAddress, connutils.RefTypeKey)

			}
		}

		// in case of error, return the error
		if err != nil {
			return err
		}

		// schema = responseSchema
		thingResponse.Schema = responseSchema

	}

	// success, return nil
	return nil

}

// thingToJanusgraph translates a thing struct into a vertex string
func (f *Janusgraph) thingToJanusgraph(UUID strfmt.UUID, objToHandle *models.Thing, addOrUpdate string) error {

	typeOfObject := "thing"

	// start vertex string Buffer
	var vertex bytes.Buffer
	var update bytes.Buffer
	var edge []string

	// add vertext class and ID
	if addOrUpdate == "add" {
		vertex.WriteString(`g.addV("` + typeOfObject + `")`)
	} else if addOrUpdate == "update" {
		vertex.WriteString(`g.V().hasLabel("` + typeOfObject + `")`)
		update.WriteString(`g.V().hasLabel("` + typeOfObject + `")`)
	} else {
		return errors_.New("addOrUpdate should be 'add' or 'update'")
	}

	// define the type and the ID
	vertex.WriteString(`.has("uuid", "` + UUID.String() + `").has("type", "` + typeOfObject + `")`)
	if addOrUpdate == "update" {
		update.WriteString(`.has("uuid", "` + UUID.String() + `").has("type", "` + typeOfObject + `").outE().drop()`)
	}

	// set the meta values (@class will be the vector)
	vertex.WriteString(`.property("atClass", "` + objToHandle.AtClass + `")`)
	vertex.WriteString(`.property("context", "` + objToHandle.AtContext + `")`)
	vertex.WriteString(`.property("creationTimeUnix", "` + strconv.FormatInt(objToHandle.CreationTimeUnix, 10) + `")`)
	vertex.WriteString(`.property("lastUpdateTimeUnix", "` + strconv.FormatInt(objToHandle.LastUpdateTimeUnix, 10) + `")`)

	// reflect in Schema
	schema := reflect.ValueOf(objToHandle.Schema)

	// fetch the schema.properties from the schema
	if schema.Kind() == reflect.Map {
		for _, e := range schema.MapKeys() {
			v := schema.MapIndex(e)
			switch t := v.Interface().(type) {
			case string:
				vertex.WriteString(`.property("schema__` + e.String() + `", "` + t + `")`)
			case int:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.Itoa(t) + `)`)
			case int8:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.Itoa(int(t)) + `)`)
			case int16:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatInt(int64(t), 10) + `)`)
			case int32:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatInt(int64(t), 10) + `)`)
			case int64:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatInt(t, 10) + `)`)
			case bool:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatBool(t) + `)`)
			case float32:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatFloat(float64(t), 'g', -1, 32) + `)`)
			case float64:
				vertex.WriteString(`.property("schema__` + e.String() + `", ` + strconv.FormatFloat(t, 'g', -1, 64) + `)`)
			case interface{}:
				if reflect.TypeOf(v.Interface()).String() == "time.Time" { // in case of time, store as date
					vertex.WriteString(`.property("schema__` + e.String() + `", "` + time.Time.String(t.(time.Time)) + `")`)
				} else if reflect.TypeOf(v.Interface()).String() == "*models.SingleRef" { // in case of SingleRef, store as relation
					// Get the singleRef values
					singleRef := v.Interface().(*models.SingleRef)
					// create the edge
					edge = append(edge, `g.addE("schema__`+e.String()+`").from(g.V().hasLabel("`+typeOfObject+`").has("uuid", uuid)).to(g.V().hasLabel("`+typeOfObject+`").has("uuid", "`+singleRef.NrDollarCref.String()+`")).property("\$cref", "`+singleRef.NrDollarCref.String()+`").property("type", "`+singleRef.Type+`").property("locationUrl", "`+*singleRef.LocationURL+`")`)
				} else {
					f.messaging.ExitError(78, "The type "+reflect.TypeOf(v.Interface()).String()+" is not found.")
				}
			default:
				f.messaging.ExitError(78, "The type "+reflect.TypeOf(v.Interface()).String()+" is not found.")
			}
		}
	}

	// execute the Vertex query with uuid as binding, result is not used because we send out "202 Accepted" and assume a succes because validation takes place before this function runs
	addResult, err := f.client.Execute(
		vertex.String(),
		map[string]string{},
		map[string]string{},
	)

	// on process error, fail
	if reflect.TypeOf(addResult.([]interface{})[0]).String() == "*errors.errorString" {
		// not returning the error because it is a go routine and the error message will arrive after the fact
		f.messaging.ErrorMessage("Janusgraph [ADD]: " + "[SCRIPT EVALUATION ERROR]")
	}

	// on error, fail
	if err != nil {
		return err
	}

	// drop all edges when updating
	if addOrUpdate == "update" {

		// execute the Edge query
		updateResult, err := f.client.Execute(
			update.String(),
			map[string]string{},
			map[string]string{},
		)

		// on process error, fail
		if reflect.TypeOf(updateResult.([]interface{})[0]) != nil {
			// not returning the error because it is a go routine and the error message will arrive after the fact
			f.messaging.ErrorMessage("Janusgraph [UPDATE]: " + "[SCRIPT EVALUATION ERROR]")
		}

		// on error, fail
		if err != nil {
			return err
		}
	}

	// if there are any edges set...
	if len(edge) > 0 {
		// loop over edges that need to be added
		for _, singleEdge := range edge {

			// execute the Edge query
			addEdgeResult, err := f.client.Execute(
				singleEdge,
				map[string]string{"uuid": UUID.String()},
				map[string]string{},
			)

			// on process error, fail
			if reflect.TypeOf(addEdgeResult.([]interface{})[0]).String() == "*errors.errorString" {
				// not returning the error because it is a go routine and the error message will arrive after the fact
				f.messaging.ErrorMessage("Janusgraph [EDGE CREATION]: " + "[SCRIPT EVALUATION ERROR]")
			}

			// on error, fail
			if err != nil {
				return err
			}
		}
	}

	// return the vertex and the edge map
	return nil
}