/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ytest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	code   int
	header http.Header
	raw    []byte
	mime   string
	body   any
}

func newResponse(resp *http.Response) *Response {
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		fatal("ReadAll resp.Body:", err)
	}
	return &Response{
		code:   resp.StatusCode,
		header: resp.Header,
		raw:    raw,
	}
}

// Code returns the status code of this response.
func (p *Response) Code() int {
	return p.code
}

func (p *Response) matchCode(t CaseT, code any) {
	t.Helper()
	switch v := code.(type) {
	case int:
		Gopt_Case_Match__0(t, p.code, v)
	case *Var__0[int]:
		v.Match(t, p.code)
	default:
		t.Fatalf("match status code failed - unexpected type: %T\n", code)
	}
}

// Header returns header of this response.
func (p *Response) Header() http.Header {
	return p.header
}

func (p *Response) matchHeader(t CaseT, key string, value any) {
	t.Helper()
	switch v := value.(type) {
	case string:
		Gopt_Case_Match__0(t, v, p.header.Get(key))
	case []string:
		Gopt_Case_Match__3(t, v, p.header[key])
	case *Var__0[string]:
		v.Match(t, p.header.Get(key))
	case *Var__3[[]string]:
		v.Match(t, p.header[key])
	default:
		t.Fatalf("match header failed! unexpected value type: %T\n", value)
	}
}

// RawBody returns raw response body as []byte (byte slice).
func (p *Response) RawBody() []byte {
	return p.raw
}

// Body returns response body decoded according Content-Type.
func (p *Response) Body() (ret any) {
	if p.mime == mimeNone {
		mime := p.header.Get("Content-Type")
		switch mime {
		case mimeJson:
			if err := json.Unmarshal(p.raw, &ret); err != nil {
				fatal("json.Unmarshal resp.Body:", err)
			}
		case mimeForm:
			form, err := url.ParseQuery(string(p.raw))
			if err != nil {
				fatal("url.ParseQuery resp.Body:", err)
			}
			ret = form
		case mimeNone:
			mime = mimeBinary
			fallthrough
		default:
			ret = p.raw
		}
		p.mime, p.body = mime, ret
	}
	return p.body
}

func (p *Response) matchBody(t CaseT, bodyType string, body any) {
	t.Helper()
	mime := p.header.Get("Content-Type")
	if mime != bodyType {
		t.Fatalf("resp.MatchBody: unmatched mime type - got: %s, expected: %s\n", mime, bodyType)
	}
	Gopt_Case_Match__4(t, body, p.Body())
}
