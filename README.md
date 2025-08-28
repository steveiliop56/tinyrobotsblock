# Tinyrobotsblock

An extremely simple (>50 lines) Traefik middleware that overwrites the `robots.txt` file and denies crawling from bots.

# Installation

To install the plugin, firstly specify it in the static configuration (`traefik.yml`) like so:

```yaml
experimental:
  plugins:
    tinyrobotsblock:
      moduleName: "github.com/steveiliop56/tinyrobotsblock"
      version: "v0.1.0"
```

Then define the middleware:

```yaml
http:
  middlewares:
    tinyrobotsblock:
      plugin:
        tinyrobotsblock:
```

# License

Tinyrobotsblock is licensed under the MIT License. TL;DR — You can use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the software. Just make sure to include the original license in any substantial portions of the code. There’s no warranty — use at your own risk. See the [LICENSE](./LICENSE) file for full details.
