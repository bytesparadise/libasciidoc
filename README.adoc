= Libasciidoc

image:https://godoc.org/github.com/bytesparadise/libasciidoc?status.svg["GoDoc", link="https://godoc.org/github.com/bytesparadise/libasciidoc"]
image:https://goreportcard.com/badge/github.com/bytesparadise/libasciidoc["Go Report Card", link="https://goreportcard.com/report/github.com/bytesparadise/libasciidoc"]
image:https://github.com/bytesparadise/libasciidoc/workflows/ci-build/badge.svg["GitHub Action Build Status", link="https://github.com/bytesparadise/libasciidoc/actions?query=workflow%3Aci-build"]
image:https://codecov.io/gh/bytesparadise/libasciidoc/branch/master/graph/badge.svg["Codecov", link="https://codecov.io/gh/bytesparadise/libasciidoc"]
image:https://img.shields.io/badge/License-Apache%202.0-blue.svg["License", link="https://opensource.org/licenses/Apache-2.0"]

Libasciidoc is an open source Go library to convert from Asciidoc to HTML.

== Supported syntax

Although it does not support the full Asciidoc/Asciidoctor syntax, Libasciidoc already provides users with the following features:

* Title and Sections level 1 to 6
* Document authors and revision
* Attribute declaration and substitution
* Paragraphs and admonition paragraphs
* Delimited Blocks (fenced blocks, listing blocks, example blocks, comment blocks, quoted blocks, sidebar blocks, verse blocks, open blocks)
* Source code highlighting of delimited blocks (use either `chroma` or `pygments` as the `source-highlighter`)
* Literal blocks (paragraph starting with a space, with the `+++....+++` delimiter or with the `[literal]` attribute)
* Quoted text (bold, italic, monospace, marked, superscript and subscript) and substitution prevention using the backslash (`\`) character
* Single and double quoted typographic quotes (e.g. '`single`' and "`double`")
* Explicit and implicit curved apostrophe
* Copyright (C), Registered (R), and Trademark (TM) symbols
* Passthrough (wrapping with a single plus or a triple plus, or using the `+++pass:[]+++` or `+++pass:q[]+++` macros)
* External links in paragraphs (`https://`, `http://`, `ftp://`, `irc://`, `mailto:`)
* Inline images in paragraphs (`image:`)
* Image blocks (`image::`)
* Icons including font, graphic icons, both in admonition blocks and inline (`icon:`)
* Element attributes (`ID`, `link`, `title`, `role`, etc.) including short-hand (`[#id.role1.role2]`)
* Ordered lists including custom numbering types (`arabic`, `upperroman`, `lowergreek`, and so forth)
* Unordered lists including bullet styles
* Labeled lists, including `[horizontal]` and `[qanda]` styles
* Nesting of links of different types & attributes
* Tables (basic support: header line and cells on multiple lines, top-level table styles)
* Horizontal rules (thematic breaks) and page breaks
* Table of contents
* YAML front-matter

See also the link:LIMITATIONS.adoc[known limitations] page for differences between Asciidoc/Asciidoctor and Libasciidoc.

Further elements will be supported in the future. Feel free to open issues https://github.com/bytesparadise/libasciidoc/issues[here] to help prioritizing the upcoming work.


=== Syntax Highlighting

When enabled, syntax highlighting in `[source]` blocks is backed by the https://github.com/alecthomas/chroma[Chroma libray]. 
The defaut class prefix is `tok-`, and it can be overridden at the document level using the `chroma-class-prefix` attribute, or from the command line interface:

[source]
----
:chroma-class-prefix: myprefix- <1>
:chroma-class-prefix: <2>
----

<1> classes with a custom prefix
<2> classes without any prefix

```
$ libasciidoc -o - -a chroma-class-prefix=myprefix- mydoc.adoc
```


== Output Formats (backend)

Using `-b` (or `--backend`) the following formats are supported:

* `html5` (also `html`), this is the default
* `xhtml5` (also `xhtml`)

== Installation

To build libasciidoc and make it available on the command line, do this:

    $ git clone https://github.com/bytesparadise/libasciidoc.git
    $ cd libasciidoc
    $ make install

If `$GOPATH/bin` is already in `$PATH`, then you should be good. Otherwise, for Linux and macOS users, you can run the following command:
 
    $ sudo ln -s "$PWD/bin/libasciidoc" /usr/local/bin/libasciidoc

== Usage

=== Command Line

The libasciidoc library includes a minimalist command line interface to generate the HTML content from a given file:

```
$ libasciidoc -s content.adoc
```

use `libasciidoc --help` to check all available options.

=== Code integration

Libasciidoc provides 2 functions to convert an Asciidoc content into HTML:

1. Converting an `io.Reader` into an HTML document:

    Convert(r io.Reader, output io.Writer, config *configuration.Configuration) (types.Metadata, error) 

2. Converting a file (given its name) into an HTML document:

    ConvertFile(output io.Writer, config *configuration.Configuration) (types.Metadata, error)

where the returned `types.Metadata` object contains the document's title which is not part of the generated HTML `<body>` part, as well as the table of contents (even if not rendered) and other attributes of the document.

All options/settings are passed via the `config` parameter.

=== Macro definition

The user can define a macro by calling `renderer.WithMacroTemplate()` and passing return value to conversion functions.

`renderer.WithMacroTemplate()` defines a macro by the given name and associates the given template. The template is an implementation of `renderer.MacroTemplate` interface (ex. `text.Template`)

Libasciidoc calls `Execute()` method and passes `types.UserMacro` object to template when rendering.

An example the following:

```
var tmplStr = `<span>Example: {{.Value}}{{.Attributes.GetAsString "suffix"}}</span>`
var t = template.New("example")
var tmpl = template.Must(t.Parse(tmplStr))

output := &strings.Builder{}
content := strings.NewReader(`example::hello world[suffix=!!!!!]`)
libasciidoc.Convert(content, output, renderer.WithMacroTemplate(tmpl.Name(), tmpl))
```

== How to contribute

Please refer to the link:CONTRIBUTE.adoc[Contribute] page.

== License

Libasciidoc is available under the terms of the https://raw.githubusercontent.com/bytesparadise/libasciidoc/LICENSE[Apache License 2.0].

== Trademark

AsciiDoc is a trademark of the Eclipse Foundation
