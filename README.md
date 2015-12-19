## Golang bindings for [libdetectcharset](https://github.com/batterseapower/libcharsetdetect.git)

This is Go bindings for the C-library [libdetectcharset](https://github.com/batterseapower/libcharsetdetect.git)
which is itself a C-wrapper for the C++ library
[Universal Charset Detector](http://mxr.mozilla.org/seamonkey/source/extensions/universalchardet/) by Mozilla.

I was interested in comparing performance of several available charset detectors
and created this bindings for this purpose.

### Performance
I bencharked four charset detectors available in Go:

* [saintfish/chardet](https://github.com/saintfish/chardet) - Go port of ICU algorithms
* [goodsign/icu](https://github.com/goodsign/icu) – libICU wrapper
* [endevit/enca](https://github.com/endeveit/enca) – libenca wrapper
* [aglyzov/charsetdetect](https://github.com/aglyzov/charsetdetect) – libdetectcharset wrapper

The test was conducted on a MacBook Air 2013 (1.3 GHz Intel Core i5) using 1 CPU
core. Test data – 4 HTML pages:

* 106 KB – UTF-8
* 105 KB – incorrect UTF-8
* 521 KB - CP1251
* 61 KB  - KOI8-R

| Detector (4x, full text) | ns/op    | B/op  | allocs/op |
| ---------------------    | -------- | ----- | --------- |
| Enca                     | 7144856  | 64    | 4         |
| CharsetDetect            | 8777701  | 64    | 4         |
| ICU                      | 36366853 | 2204  | 60        |
| Chardet                  | 68107438 | 53828 | 32        |

| Detector (4x, 4096 bytes) | ns/op   | B/op  | allocs/op |
| ------------------------- | ------- | ----- | --------- |
| Enca                      | 89757   | 64    | 4         |
| CharsetDetect             | 187303  | 64    | 4         |
| ICU                       | 4393345 | 53824 | 32        |
| Chardet                   | 8906792 | 2165  | 60        |


[endevit/enca](https://github.com/endeveit/enca) was a clear winner both in speed and quality.
[aglyzov/charsetdetect](https://github.com/aglyzov/charsetdetect) was coming the second.
So personally I recommend using [endevit/enca](https://github.com/endeveit/enca) instead.

