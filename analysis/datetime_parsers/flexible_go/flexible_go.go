//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package flexible_go

import (
	"fmt"
	"time"

	"github.com/tukdesk/bleve/analysis"
	"github.com/tukdesk/bleve/registry"
)

const Name = "flexiblego"

type FlexibleGoDateTimeParser struct {
	layouts []string
}

func NewFlexibleGoDateTimeParser(layouts []string) *FlexibleGoDateTimeParser {
	return &FlexibleGoDateTimeParser{
		layouts: layouts,
	}
}

func (p *FlexibleGoDateTimeParser) ParseDateTime(input string) (time.Time, error) {
	for _, layout := range p.layouts {
		rv, err := time.Parse(layout, input)
		if err == nil {
			return rv, nil
		}
	}
	return time.Time{}, analysis.ErrInvalidDateTime
}

func FlexibleGoDateTimeParserConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.DateTimeParser, error) {
	layouts, ok := config["layouts"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("must specify layouts")
	}
	layoutStrs := make([]string, 0)
	for _, layout := range layouts {
		layoutStr, ok := layout.(string)
		if ok {
			layoutStrs = append(layoutStrs, layoutStr)
		}
	}
	return NewFlexibleGoDateTimeParser(layoutStrs), nil
}

func init() {
	registry.RegisterDateTimeParser(Name, FlexibleGoDateTimeParserConstructor)
}
