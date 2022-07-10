# ESM

A fast, global content delivery network for [NPM](http://npmjs.org/) packages with **ES Module** format.

## Import from URL

```javascript
import React from "https://esm.sh/react"
```

### Specify version

```javascript
import React from "https://esm.sh/react@17.0.2"
```

You may also use a [semver](https://docs.npmjs.com/cli/v6/using-npm/semver) or a [dist-tag](https://docs.npmjs.com/cli/v8/commands/npm-dist-tag) instead of a fixed version number, or omit the version/tag entirely to use the `latest` tag:

```javascript
import React from "https://esm.sh/react@17"   // 17.0.2
import React from "https://esm.sh/react@next" // 18.0.0-rc.0-next-13036bfbc-20220121
```

### Submodule

```javascript
import { renderToString } from "https://esm.sh/react-dom/server"
```

or import non-module(js) files:

```javascript
import "https://esm.sh/react/package.json" assert { type: "json" }
```

You can also use the `?path` to specify the `submodule`, this is friendly for **import maps**:

```json
{
  "imports": {
    "react-dom/": "https://esm.sh/react-dom?target=es2015&path=/"
  }
}
```

```javascript
// equals to https://esm.sh/react-dom/server?target=es2015
import { renderToString } from "react-dom/server"
```

### Bundle mode

```javascript
import { Button } from "https://esm.sh/antd?bundle"
```

In **bundle** mode, all dependencies will be bundled into a single JS file.

### Development mode

```javascript
import React from "https://esm.sh/react?dev"
```

The `?dev` mode builds modules with `process.env.NODE_ENV` equals to `development`, that is useful to build modules like **React** to allow you to get more development warn/error details.

### Specify external dependencies

```javascript
import React from "https://esm.sh/react@16.14.0"
import useSWR from "https://esm.sh/swr?deps=react@16.14.0"
```

By default, esm.sh rewrites import specifier based on the package's dependency statement. To specify version of dependencies, you can use the `?deps=PACKAGE@VERSION` query. You can separate multiple dependencies with commas: `?deps=react@16.14.0,react-dom@16.14.0`.

### Aliasing dependencies

```javascript
import useSWR from "https://esm.sh/swr?alias=react:preact/compat"
```

in combination with `?deps`:

```javascript
import useSWR from "https://esm.sh/swr?alias=react:preact/compat&deps=preact@10.5.14"
```

The origin idea was coming from [@lucacasonato](https://github.com/lucacasonato).

### ESBuild options

By default, esm.sh will check the `User-Agent` header to get the build target automatically. You can specify it with the `?target` query. Available targets: **es2015** - **es2022**, **esnext**, **node**, and **deno**.

```javascript
import React from "https://esm.sh/react?target=es2020"
```

Other supported options of [esbuild](https://esbuild.github.io/):

- [Keep names](https://esbuild.github.io/api/#keep-names)
  ```javascript
  import React from "https://esm.sh/react?keep-names"
  ```
- [Ignore annotations](https://esbuild.github.io/api/#ignore-annotations)
  ```javascript
  import React from "https://esm.sh/react?ignore-annotations"
  ```

### Package CSS

```javascript
import Daygrid from "https://esm.sh/@fullcalendar/daygrid"
```

```html
<link rel="stylesheet" href="https://esm.sh/@fullcalendar/daygrid?css">
```

This only works when the NPM module imports CSS files in JS directly.


## Web Worker

esm.sh supports `?worker` mode to load modules as web worker:

```javascript
import editorWorker from "https://esm.sh/monaco-editor/esm/vs/editor/editor.worker?worker"

const worker = editorWorker()
```

## Deno compatibility

**esm.sh** will resolve the node internal modules (**fs**, **child_process**, etc.) with [`deno.land/std/node`](https://deno.land/std/node) to support some packages working in Deno, like `postcss`:

```javascript
import postcss from "https://esm.sh/postcss"
import autoprefixer from "https://esm.sh/autoprefixer"

const { css } = await postcss([ autoprefixer ]).process(`
  backdrop-filter: blur(5px);
  user-select: none;
`).async()
```

### X-Typescript-Types

By default, **esm.sh** will respond with a custom `X-TypeScript-Types` HTTP header when the types (`.d.ts`) is defined. This is useful for deno type checks ([link](https://deno.land/manual/typescript/types#using-x-typescript-types-header)).

![Figure #1](./server/embed/assets/sceenshot-deno-types.png)

You can pass the `?no-dts` query to disable the `X-TypeScript-Types` header if some types are incorrect:

```javascript
import unescape from "https://esm.sh/lodash/unescape?no-dts"
```

## Pin the build version

Since we update esm.sh server frequently, sometime we may break packages that work fine previously by mistake, the server will rebuild all modules when the patch pushed. To avoid this, you can **pin** the build version by the `?pin=BUILD_VERSON` query. This will give you an **immutable** cached module.

```javascript
import React from "https://esm.sh/react@17.0.2?pin=v86"
```

## Global CDN

<img width="150" align="right" src="./server/embed/assets/cf.svg">

The Global CDN of esm.sh is provided by [Cloudflare](https://cloudflare.com), one of the world's largest and fastest cloud network platforms.

## Self-Hosting

To host esm.sh by yourself, check the [hosting](./HOSTING.md) documentation.
