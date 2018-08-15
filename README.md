# jameslucktaylor.info

My Go-powered website!

Hosted on [Google Cloud](https://cloud.google.com)'s [App Engine](https://cloud.google.com/appengine/) under the [Go Standard Environment](https://cloud.google.com/appengine/docs/standard/go/).

I have been tinkering with this site as a vehicle to learn some more web dev things, and am trying to put best practices and good security and industry standards and the like into place as I go.

The page itself doesn't have very much functionality at all, just some links to me elsewhere on the web, and the real value in this project for me is (over-)engineering everything behind and around it.

## Checks

[![SSL Rating](https://sslbadge.org/?domain=jameslucktaylor.info)](https://ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info)
[![Why No Padlock?](https://img.shields.io/badge/Why%20No%20Padlock%3F-Pass-brightgreen.svg?style=plastic)](https://whynopadlock.com/results/c80ada01-1136-4321-9819-efab5b6c3205)

## Tools used

### Build

- [Go](https://golang.org)
- [Google Cloud SDK](https://cloud.google.com/sdk/)

### Design

- [Font Awesome](https://fontawesome.com)
  - [SVG & JS](https://fontawesome.com/how-to-use/on-the-web/setup/getting-started?using=svg-with-js)
- [Roboto - Google Fonts](https://fonts.google.com/specimen/Roboto)

### Validate

#### Security

- [OWASP Zed Attack Proxy](https://owasp.org/index.php/OWASP_Zed_Attack_Proxy_Project)
- [Qualys SSL Labs - SSL Server Test](https://ssllabs.com/ssltest/)
- [Why No Padlock?](https://whynopadlock.com)

#### Performance

- [Google PageSpeed Insights](https://developers.google.com/speed/pagespeed/insights/)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse/)

#### Responsiveness

- [hey](https://github.com/rakyll/hey)

#### Functionality

- [Favicon Checker](https://realfavicongenerator.net/favicon_checker)
- Facebook for Developers
  - [Sharing Debugger](https://developers.facebook.com/tools/debug/sharing/)
  - [Object Debugger](https://developers.facebook.com/tools/debug/og/object/)
- [Google Developers - Open Graph structured data](https://developers.google.com/search/docs/guides/prototype)
- [W3C Validator](http://validator.w3.org)

### Miscellaneous (blogs, articles, references)

I've proud/ashamed to admit that I have chased a lot of rabbits 🐇 down their burroughs and [shaved quite a few yaks](https://www.youtube.com/watch?v=AbSehcT19u0) while working on this project. 😅

- [Lighthouse: how to reduce render-blocking scripts](https://fly.io/articles/lighthouse-how-to-reduce-render-blocking-scripts/)
- Font-display rabbit hole (thanks to Lighthouse for highlighting this)
  - [A small explainer built for a talk on web fonts and performance](https://font-display.glitch.me)
  - [If you really dislike FOUT, `font-display: optional` might be your jam](https://css-tricks.com/really-dislike-fout-font-display-optional-might-jam/)
  - [A comprehensive guide to font loading strategies](https://www.zachleat.com/web/comprehensive-webfonts/#font-display)

## TODO

- [Terraform](https://terraform.io) the Google Cloud infrastructure behind the site
- Add the [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) header
- Configure [warmup requests](https://cloud.google.com/appengine/docs/standard/go/warmup-requests/configuring) to improve performance
- Slim down Font Awesome to just a few characters, rather than the whole library
  - [Performance & Font Awesome](https://fontawesome.com/how-to-use/on-the-web/other-topics/performance)
- Set up a [CI/CD pipeline with Cloud Build](https://cloud.google.com/community/tutorials/automated-publishing-cloud-build)
  - Roll some [automated testing](https://cloud.google.com/cloud-build/docs/configuring-builds/build-test-deploy-artifacts) into said pipeline
- Configure [a CDN](https://cloud.google.com/cdn/docs/using-cdn) (and the [load balancing](https://cloud.google.com/load-balancing/docs/https/) pre-requisite) on Google Cloud
- Finalise [HSTS preload](https://hstspreload.org/)
- Send OWASP ZAP output through [a parser](https://yq.readthedocs.io/en/latest/) and make it more succinct
- Optimise the GitHub fork image with one of the following methods:
  - [Convert the image file itself](https://developers.google.com/speed/docs/insights/OptimizeImages)
  - [Use these nice third-party SVGs](https://github.com/tholman/github-corners)
