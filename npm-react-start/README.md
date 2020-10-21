# npm-react-start

A simple filesystem benchmark which initially bootstraps a React app with
`npx create-react-app` and then runs `npm start`, timing how long it takes
for the development webserver to start.

## Usage:

```
docker run -v /Users/djs/workspace:/volume djs55/npm-react-start
```

The output is in .csv format:
```
# npx create-react-app => npm start benchmark
# app has been bootstrapped
# iteration, time/seconds
#
# > my-app@0.1.0 start /volume/my-app
# > react-scripts start
#
# ℹ ｢wds｣: Project is running at http://172.17.0.2/
# ℹ ｢wds｣: webpack output is served from
# ℹ ｢wds｣: Content not from webpack is served from /volume/my-app/public
# ℹ ｢wds｣: 404s will fallback to /
# Starting the development server...
#
# Browserslist: caniuse-lite is outdated. Please run the following command: `npx browserslist --update-db`
# Compiled successfully!
0, 2.306320
...
```