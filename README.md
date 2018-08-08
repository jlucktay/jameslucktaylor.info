# jameslucktaylor.info

My Go-powered website!

Hosted on [Google Cloud](https://cloud.google.com)'s [App Engine](https://cloud.google.com/appengine/) under the [Go Standard Environment](https://cloud.google.com/appengine/docs/standard/go/).

## Checks

[![SSL Rating](https://sslbadge.org/?domain=jameslucktaylor.info)](https://www.ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info)

## Tools used

### Build

- [Go](https://golang.org)
- [Google Cloud SDK](https://cloud.google.com/sdk/)

### Design

- [Roboto - Google Fonts](https://fonts.google.com/specimen/Roboto)
- [Font Awesome](https://fontawesome.com)
  - [SVG & JS](https://fontawesome.com/how-to-use/on-the-web/setup/getting-started?using=svg-with-js)

### Validate

#### Security

- [OWASP Zed Attack Proxy](https://www.owasp.org/index.php/OWASP_Zed_Attack_Proxy_Project)
- [Qualys SSL Labs - SSL Server Test](https://www.ssllabs.com/ssltest/)
- [Why No Padlock?](https://www.whynopadlock.com)

#### Functionality

- [Favicon Checker](https://realfavicongenerator.net/favicon_checker)
- [hey](https://github.com/rakyll/hey)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse/)
  - [Lighthouse: how to reduce render-blocking scripts](https://fly.io/articles/lighthouse-how-to-reduce-render-blocking-scripts/)
- [W3C Validator](http://validator.w3.org)

## TODO

- Add the [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) header
- Configure [warmup requests](https://cloud.google.com/appengine/docs/standard/go/warmup-requests/configuring) to improve performance
- Slim down Font Awesome to just a few characters, rather than the whole library
  - [Performance & Font Awesome](https://fontawesome.com/how-to-use/on-the-web/other-topics/performance)
- Set up a [CI/CD pipeline with Cloud Build](https://cloud.google.com/community/tutorials/automated-publishing-cloud-build)
  - Roll some [automated testing](https://cloud.google.com/cloud-build/docs/configuring-builds/build-test-deploy-artifacts) into said pipeline
- Configure [a CDN](https://cloud.google.com/cdn/docs/using-cdn) (and the [load balancing](https://cloud.google.com/load-balancing/docs/https/) pre-req) on Google Cloud
- Finalise [HSTS preload](https://hstspreload.org/)
