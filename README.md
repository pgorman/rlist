Rlist
========================================

Rlist is a simple text filter to format ranged lists in Markdown and plain text into HTML or LaTeX.

Ranged lists are common in some domains, like role playing games. For example, roll an eight-sided die to determine magic fruit:

```
1. Poison apple: save or die
2–3. Banana: heals 1 hit point
4–7. Orange: removes any temporary minor adverse condition (-1 penalties)
8. Golden apple: permanently increases random ability score by one
```

Unfortunately, none of the markup languages include semantic ranged lists.
The closest we get is definition lists.
From the example input above, Rlist produces this HTML output:

```html
<dl class="rlist">
<dt>1.</dt>
<dd>Poison apple: save or die</dd>
<dt>2–3.</dt>
<dd>Banana: heals 1 hit point</dd>
<dt>4–7.</dt>
<dd>Orange: removes any temporary minor adverse condition (-1 penalties)</dd>
<dt>8.</dt>
<dd>Golden apple: perminently increases random ability score by one</dd>
</dl>
```

Then do whatever CSS we want, like:

```css
dl.rlist > dt {
	clear: left;
	float: left;
	margin: 0;
	margin-right: 1em;
	padding: 0;
	text-align: center;
	width: 3em;
}
dl.rlist > dd {
	margin-left: 4em;
}
```

Use Rlist in a command pipeline, like a standard Unix filter.

```
$  rlist < mytext.md
```

Commonly, the next stage in the pipeline will be a Markdown processor.

```
$  rlist < mytext.md | cmark
```

By default, Rlist outputs HTML. Use the `-x` flag for LaTeX output, or `-h` to see the help for all flags.


Limitations
---------------------------------------

Rlist considers a list to have ended on the first non-blank line that does not begin with a number. Consequently, it does not support multi-paragraph list items.

Because Rlist needs to look ahead to see where a list ends and its type, it reads the entire list into memory. Very, very large lists might fail.


License (2-Clause BSD)
---------------------------------------

Copyright 2019 Paul Gorman

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.