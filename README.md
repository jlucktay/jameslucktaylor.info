# jameslucktaylor.info

My Go-powered website!

Hosted on [Google Cloud](https://cloud.google.com)'s [App Engine](https://cloud.google.com/appengine/) under the [Go Standard Environment](https://cloud.google.com/appengine/docs/standard/go/).

## Tools used

### Build

- [Go](https://golang.org)
- [Google Cloud SDK](https://cloud.google.com/sdk/)

### Validate

#### Security

- [OWASP Zed Attack Proxy](https://www.owasp.org/index.php/OWASP_Zed_Attack_Proxy_Project)
- [Qualys SSL Labs - SSL Server Test](https://www.ssllabs.com/ssltest/)
- [Why No Padlock?](https://www.whynopadlock.com)

#### Functionality

- [Favicon Checker](https://realfavicongenerator.net/favicon_checker)
- [hey](https://github.com/rakyll/hey)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse/)
- [W3C Validator](http://validator.w3.org)

## TODO

- Add the [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) header
- Configure [warmup requests](https://cloud.google.com/appengine/docs/standard/go/warmup-requests/configuring) to improve performance
