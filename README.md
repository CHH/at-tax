at-tax
======

Simple web service which estimates the Austrian income tax for
self employed people like me (not for employed).

It only returns an estimate, without warranty!

## Usage

Make a HTTP POST request to `/income-tax` with a JSON payload. The payload's
format looks like this:

```json
{"income": 14000.21}
```

`income` is the amount in Euro as a floating point number.

The response format looks like this:

```json
{"tax": 1095.08}
```

## Errors

HTTP error codes are used throughout the API. A detailed error message is
returned in the JSON's `reason` property.

## License

The MIT License (MIT)
Copyright © 2014 Christoph Hochstrasser <me@christophh.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
