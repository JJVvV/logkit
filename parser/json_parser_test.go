package parser

import (
	"encoding/json"
	"testing"

	"github.com/qiniu/logkit/conf"
	"github.com/qiniu/logkit/sender"
	"github.com/qiniu/logkit/utils"

	"bytes"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestJsonParser(t *testing.T) {
	c := conf.MapConf{}
	c[KeyParserName] = "testjsonparser"
	c[KeyParserType] = "json"
	c[KeyLabels] = "mm abc"
	p, _ := NewJsonParser(c)
	tests := []struct {
		in  []string
		exp []sender.Data
	}{
		{
			in: []string{`{"a":1,"b":[1.0,2.0,3.0],"c":{"d":"123","g":1.2},"e":"x","f":1.23}`},
			exp: []sender.Data{sender.Data{
				"a": json.Number("1"),
				"b": []interface{}{json.Number("1.0"), json.Number("2.0"), json.Number("3.0")},
				"c": map[string]interface{}{
					"d": "123",
					"g": json.Number("1.2"),
				},
				"e":  "x",
				"f":  json.Number("1.23"),
				"mm": "abc",
			}},
		},
		{
			in: []string{`{"a":1,"b":[1.0,2.0,3.0],"c":{"d":"123","g":1.2},"e":"x","mm":1.23,"jjj":1493797500346428926}`},
			exp: []sender.Data{sender.Data{
				"a": json.Number("1"),
				"b": []interface{}{json.Number("1.0"), json.Number("2.0"), json.Number("3.0")},
				"c": map[string]interface{}{
					"d": "123",
					"g": json.Number("1.2"),
				},
				"e":   "x",
				"mm":  json.Number("1.23"),
				"jjj": json.Number("1493797500346428926"),
			}},
		},
	}
	for _, ti := range tests {
		m, err := p.Parse(ti.in)
		if err != nil {
			errx, _ := err.(*utils.StatsError)
			if errx.ErrorDetail != nil {
				t.Error(errx.ErrorDetail)
			}
		}
		assert.EqualValues(t, ti.exp, m)
	}
	assert.EqualValues(t, "testjsonparser", p.Name())
}

var testjsonline = `{"a":1,"b":[1.0,2.0,3.0],"c":{"d":"123","g":1.2},"e":"x","mm":1.23,"jjj":1493797500346428926}`
var testmiddleline = `{
  "person": {
    "id": "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "name": {
      "fullName": "Leonid Bugaev",
      "givenName": "Leonid",
      "familyName": "Bugaev"
    },
    "email": "leonsbox@gmail.com",
    "gender": "male",
    "location": "Saint Petersburg, Saint Petersburg, RU",
    "geo": {
      "city": "Saint Petersburg",
      "state": "Saint Petersburg",
      "country": "Russia",
      "lat": 59.9342802,
      "lng": 30.3350986
    },
    "bio": "Senior engineer at Granify.com",
    "site": "http://flickfaver.com",
    "avatar": "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "employment": {
      "name": "www.latera.ru",
      "title": "Software Engineer",
      "domain": "gmail.com"
    },
    "facebook": {
      "handle": "leonid.bugaev"
    },
    "github": {
      "handle": "buger",
      "id": 14009,
      "avatar": "https://avatars.githubusercontent.com/u/14009?v=3",
      "company": "Granify",
      "blog": "http://leonsbox.com",
      "followers": 95,
      "following": 10
    },
    "twitter": {
      "handle": "flickfaver",
      "id": 77004410,
      "bio": null,
      "followers": 2,
      "following": 1,
      "statuses": 5,
      "favorites": 0,
      "location": "",
      "site": "http://flickfaver.com",
      "avatar": null
    },
    "linkedin": {
      "handle": "in/leonidbugaev"
    },
    "googleplus": {
      "handle": null
    },
    "angellist": {
      "handle": "leonid-bugaev",
      "id": 61541,
      "bio": "Senior engineer at Granify.com",
      "blog": "http://buger.github.com",
      "site": "http://buger.github.com",
      "followers": 41,
      "avatar": "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
    },
    "klout": {
      "handle": null,
      "score": null
    },
    "foursquare": {
      "handle": null
    },
    "aboutme": {
      "handle": "leonid.bugaev",
      "bio": null,
      "avatar": null
    },
    "gravatar": {
      "handle": "buger",
      "urls": [
      ],
      "avatar": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
      "avatars": [
        {
          "url": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
          "type": "thumbnail"
        }
      ]
    },
    "fuzzy": false
  },
  "company": null
}`

//BenchmarkJsoninterParser-4                       	  300000	      5144 ns/op
func BenchmarkJsoninterParser(b *testing.B) {
	jsonnumber := jsoniter.Config{
		EscapeHTML: true,
		UseNumber:  true,
	}.Froze()
	for i := 0; i < b.N; i++ {
		data := sender.Data{}
		if err := jsonnumber.Unmarshal([]byte(testjsonline), &data); err != nil {
			b.Error(err)
		}
	}
}

//BenchmarkJsonParser-4                    	  200000	      7767 ns/op
func BenchmarkJsonParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := sender.Data{}
		decoder := json.NewDecoder(bytes.NewReader([]byte(testjsonline)))
		decoder.UseNumber()
		if err := decoder.Decode(&data); err != nil {
			b.Error(err)
		}
	}
}

//BenchmarkJsonMiddlelineParser-4                  	   30000	     58441 ns/op
func BenchmarkJsonMiddlelineParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := sender.Data{}
		decoder := json.NewDecoder(bytes.NewReader([]byte(testmiddleline)))
		decoder.UseNumber()
		if err := decoder.Decode(&data); err != nil {
			b.Error(err)
		}
	}
}

//BenchmarkJsoniterMiddlelineWithDecoderParser-4   	   30000	     41496 ns/op
func BenchmarkJsoniterMiddlelineWithDecoderParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := sender.Data{}
		decoder := jsoniter.NewDecoder(bytes.NewReader([]byte(testmiddleline)))
		decoder.UseNumber()
		if err := decoder.Decode(&data); err != nil {
			b.Error(err)
		}
	}
}

//BenchmarkMiddlelineWithConfigParser-4            	   50000	     35298 ns/op
func BenchmarkMiddlelineWithConfigParser(b *testing.B) {
	jsonnumber := jsoniter.Config{
		EscapeHTML: true,
		UseNumber:  true,
	}.Froze()
	for i := 0; i < b.N; i++ {
		data := sender.Data{}
		if err := jsonnumber.Unmarshal([]byte(testmiddleline), &data); err != nil {
			b.Error(err)
		}
	}
}
