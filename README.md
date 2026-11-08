# My opinionated static gallery generator

[![build badge](https://github.com/theobori/recueil/actions/workflows/build.yml/badge.svg)](https://github.com/theobori/recueil/actions/workflows/build.yml)

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

The name of the project is "recueil", pronounced \\ʁə.kœj\\.

This GitHub repository is a [KISS](https://en.wikipedia.org/wiki/KISS_principle) project that contains a CLI wrote in the Go language. The goal of this CLI is to build an image gallery in the form of a simple, minimalist static website without JavaScript. This project is directly inspired by [Ocular](https://codeberg.org/rostiger/ocular/) and [img.corvid.cafe](https://img.corvid.cafe/).

## Getting started

Before continuing to read the documentation, make sure you have at least [Go](https://go.dev/dl/) version 1.26.10.

### Installation

If you want to use the CLI, you can build and install the binary using the following command.

```bash
go install github.com/theobori/recueil/cmd/recueil
```

To start using the tool, simply run the following command.

```bash
recueil -help
```

## Specifications

This section contains the project specifications.

### Source directory file structure

In the source directory, the program retrieves only the files at depth 0, that is, at the root of the directory.

### Image Metadata

Images can optionally include metadata. This metadata must be written to a file in YAML format and placed next to the image file. The name of this YAML file must consist of the image name followed by `.yaml`. For example, for a file named `my-image.png`, the metadata file would be named `my-image.png.yaml`.

The table below summarizes the YAML configuration specifications.

| Field | Description | Field type | Default value |
| - | - | - | - |
| `tags` | The image tags. | List of string | `[]` |
| `description` | The image description. | String | `""` |
| `title` | The image title. | String | The image filename. |

### Configuration

The table below summarizes the YAML configuration specifications.

| Field | Description | Field type | Default value |
| - | - | - | - |
| `baseDirectory` | The base directory within the destination directory. It contains every image and the associated HTML pages. | String | "/recueil" |
| `title` | The website title. | String | `"My recueil"` |
| `css.colorBg` | The website background color. | String | `"#fffef7"` |
| `css.colorFg` | The website foreground color. | String | `"#394655"` |
| `css.colorLink` | The website links color. | String | `"brown"` |
| `css.colorHover` | The website hovers color. | String | `"brown"` |
| `css.colorSelectionFg` | The website text selection color. | String | `"#fffef7"` |
| `css.colorSelectionBg` | The website text selection background color. | String | `"brown"` |
| `siteUrl` | The website URL. | String | `"http://localhost:8080"` |
| `publicSourceDirectory` | The source public directory. The files inside will be available from the website root directory. As example, You can use it to add a favicon file. | String | `"./public"` |
| `images.maxPixelSize` | The images max pixel size. If images.autoResize is set to true, it will resize images with a width or a height greater than this value. | Integer | `960` |
| `images.autoResize` | Enable the images autoresize feature. | Boolean | `true` |
| `thumbnailPixelSize` | The thumbnails size, used for width and height. | Integer | `256` |
| `webringUrl` | A single webring URL (It should probably be a list). | String | `""` |

Below is a YAML configuration with all these fields populated with their default values.

```yaml
baseDirectory: "/recueil"
title: "my recueil"
css:
  colorBg: "#fffef7"
  colorFg: "#394655"
  colorLink: "brown"
  colorHover: "brown"
  colorSelectionFg: "#fffef7"
  colorSelectionBg: "brown"
siteUrl: "http://localhost:8080"
publicSourceDirectory: "./public"
images:
  maxPixelSize: 960
  autoResize: true
thumbnailPixelSize: 256
webringUrl: ""
```

### RSS feed

The project generates an XML file containing an RSS feed that complies with the standards described by the [Ocular Webring](https://codeberg.org/rostiger/ocular_webring).

## Customization

If you want to customize the style of the web pages, I think the easiest way is to override the `gallery.css` and `image.css` files in the root directory of the built website. To do this, you can create files with the same names in your public source folder; its path is specified in the YAML configuration via the `publicSourceDirectory` field.

## Example

You'll find a sample project for a collection in the [example](./example/) folder.

## Live Demo

You can find an example of a static gallery built with the tool and deployed at https://recueil-demo.theobori.cafe.

## Contribute

If you'd like to contribute to the project, please follow the instructions provided in the [CONTRIBUTING.md](./CONTRIBUTING.md) file.
