# HackerRank: Attribute Parser

## Specs
- Each element consists of a starting and an ending tag
- There are attributes assoc with each tag
- Only starting tags can have attrib
- We can call an attr by ref'ing a tag followed by a `~' and the attr name
- Tags can be nested

## Opening tags fmt:
  <tag-name attribute1-name = "value1" attribute2-name = "value2" ...>

## Closing tags fmt:
  </tag-name>

## Attr are ref'ed as:
  tag1~value
  tag1.tag2~name

### Input
The source consisting of N lines answer Q queries, for each query print
the value of the attr specified. pritn "Not Found!" if the attr doesnt exist.
