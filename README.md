# Pantoc
Parse a table of content to build a splitted document in mardown. Nice tool for pandoc

# Installation

Download released binary at [on github](https://github.com/metal3d/pantoc/releases).

Or use "go get -u":

    $ go get -u github.com/metal3d/pantoc
    $ pantoc -h


# Usage

```
Usage of pantoc:
  -tocfile="toc.yaml": YAML file that describes table of contents
```

# Yaml format

Create a yaml file inside your document tree. The Yaml file **should** define a list ! That means that you have to prefix document part with a "minus" sign.

Example

```yaml
- Introduction: src/intro.md
- Title for chapter 2: src/part2.md #a only one file for that chapter
- Title for other part:
    - src/intro_part_3 # this will be included without title
    - Sub part: src/subpart1.md # title is "Sub part" and file is included
    - Sub part2: src/subpart2.md # as above
```

Each key becomes the title, each file path are appended after title.

You can now try to build your document:

    pantoc

If you named your tocfile with a different name:

    pantoc -tocfile=mytoc.yaml

The command should print the entire document to STDOUT (your terminal).

Use it with [pandoc](http://johnmacfarlane.net/pandoc/):

    # if you have toc.yaml
    pandoc <(pantoc) -o book.pdf
    
    # with another file
    pandoc <(pantoc -tocfile=mytoc.yaml) -o book.pdf

You may have the book.pdf file that respect you table of content and with included content.


