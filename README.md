# strand

[![GoDoc](https://godoc.org/github.com/Praseetha-KR/strand?status.svg)](https://godoc.org/github.com/Praseetha-KR/strand)


#### A tiny go package for generating random strings

`strand` generates string randoms containing alphabets, digits and special characters. The randomness is seeded internally with the timestamp of package initialization.

### Install

```bash
go get -u github.com/Praseetha-KR/strand
```

### Sample Usage:

```go
package main

import (
  "fmt"
  "github.com/Praseetha-KR/strand"
)

func main() {
  fmt.Println(strand.String(24))
}
```

### Available Methods

<table>
    <thead>
        <tr>
            <th>Method</th>
            <th>Character Set</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <h5><code>String(length)</code></h5>
            </td>
            <td>All supported characters</td>
            <td>
                <div>Usage: <code>r := strand.String(20)</code></div>
                <div>Output: <code>vT7+Bt/])o+66H<[GG[s</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>Alpha(length)</code></h5>
            </td>
            <td>Alphabet</td>
            <td>
                <div>Usage: <code>r := strand.Alpha(20)</code></div>
                <div>Output: <code>PNsxdCsGyLDEFzQxcZSy</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>AlphaLower(length)</code></h5>
            </td>
            <td>Alphabet Lowercase</td>
            <td>
                <div>Usage: <code>r := strand.AlphaLower(20)</code></div>
                <div>Output: <code>ivkosjouadiecjvdanct</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>AlphaUpper(length)</code></h5>
            </td>
            <td>Alphabet Uppercase</td>
            <td>
                <div>Usage: <code>r := strand.AlphaUpper(20)</code></div>
                <div>Output: <code>QUXMYZJUHIOZPOAAASFI</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>AlphaNumeric(length)</code></h5>
            </td>
            <td>Alphabet and Digits</td>
            <td>
                <div>Usage: <code>r := strand.AlphaNumeric(20)</code></div>
                <div>Output: <code>1CxlOZp3yVwoGeqIjimH</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>Numeric(length)</code></h5>
            </td>
            <td>Digits</td>
            <td>
                <div>Usage: <code>r := strand.Numeric(20)</code></div>
                <div>Output: <code>93422189029287851575</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>URLSafe(length)</code></h5>
            </td>
            <td>url-safe characters</td>
            <td>
                <div>Usage: <code>r := strand.URLSafe(20)</code></div>
                <div>Output: <code>P~8i1aEP_PGWXyHUVmO8</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>Hex(length)</code></h5>
            </td>
            <td>Hexadecimal characters</td>
            <td>
                <div>Usage: <code>r := strand.Hex(20)</code></div>
                <div>Output: <code>b0590721d92ce097b8a3</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>Binary(length)</code></h5>
            </td>
            <td>Binary digits</td>
            <td>
                <div>Usage: <code>r := strand.Binary(20)</code></div>
                <div>Output: <code>10100011000101110011</code></div>
            </td>
        </tr>
        <tr>
            <td>
                <h5><code>From(characters, length)</code></h5>
            </td>
            <td>Provided by `characters` param</td>
            <td>
                <div>Usage: <code>r, err := strand.From("abc#123", 20)</code></div>
                <div>Output: <code>#1a#cc33ccb2131a1233</code></div>
            </td>
        </tr>
    </tbody>
</table>


---

Author: Praseetha-KR ([@void_imagineer](https://twitter.com/void_imagineer))

License: [MIT](LICENSE)
