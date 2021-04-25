# Frongo
Frongo is a Go tool to make HTML/CSS document out of Golang code.
It was designed with readability and usability in mind, so HTML objects are
created by chaining method calls.

Quick example :
```
div := *frongo.Div()
```

Initializes a div to be inserted into the document later on.

```
div.Div().Class("class").Attr("id", "testId").Write("some text")
```

Creates a div tag inside the div object, and adds a class, 
an attribute and some text into it.

But you can also get the same result by using :

```
div.Div("class", "id", "testId").Write("some text")
```

However the first way is recommended since it optimizes readability.

Either way, You will obtain this : 

```
<div>
    <div class="class" id="testId">
        some text
    </div>
</div>
```

# Installation

Like most Go packages, the installation is extremely simple :
```
$ go get -u github.com/eleby/pixelizer
```

And now you can use Frongo in your project. 

# Usage

Frongo reproduces the standard structure for a web document :

Document\
----Containers\
--------Tags\
----Sheets\
--------Styles

So in Frongo, the first thing to do is to create a document.

```
doc := frongo.InitDocument()
```

Then we want to create the objects we will use in our document.

```
body := *frongo.InitContainer("body")
sheet := frongo.Sheet{}
```
Once we have done these initializations, we can get to coding
our html and css elements.

```
div := *frongo.NewTag("div")
```

The "NewTag" function can be replaced by presets to improve 
readability as well as usability :


```
title := *frongo.Title("title text")
div := *frongo.Div()
span := *frongo.Span()
paragraph := *frongo.P()
img := *frongo.Img()
lineJump := *frongo.Br()
```

All these presets can also be used when adding a tag's child, like 
div.Div() which adds a div to the first div.

All presets except Title and Br use the "params" parameter which indicates 
a potential class and list of attributes, attribute values and delimiter between
each attribute name/value. All these are optional, and are shown (except the delimiter
which is "," by default) in the next example. The Title preset uses a mandatory string
as the text to use as the title.

So once we have our div, we can easily add classes and attributes;
there are many ways to do this. 

At initialization :
```
div := *frongo.Div("class", "attr1,attr2", "val1,val2")
```

Later, all at once :

```
div.Class("class").Attrs("attr1,attr2", "val1,val2")
```

Or later, one by one :

```
div.Class("class").Attr("attr1", "val1").Attr("attr2", "val2")
```

You can also make a mix :

```
div := *frongo.Div("class").Attr("attr1", "val1").Attr("attr2", "val2")
```

Having all these possibilities allow you to manage your code's 
length and readability so you can use it the way which fits you.

If you want to add text to a div, you can use the Write method 
to write IN the div (before its children) or WriteAfter to write
after the div. 


```
div.Write("some text") //<div>some text</div>
```
```
div.WriteAfter("some text") //<div></div>some text
```

Now if you want to close the tag (so it will be a self-closing one)
you can use the Close method.

```
linkTag.Close()
```

So now we have seen the basics for HTML tags. There are also many
methods to find them by tag name or by id, or index :

```
div.FindById("testId") //Gets the first found element with id
div.Find("p") //Gets an array of all the children p tags
div.First() //Gets the first child of the div
div.Get(1) //Gets the second child of the div
div.Last() //Gets the last child of the div
```

The Find method also can be called on the document object once
the tags are added. 

There are also check methods :

```
document.Contains("div") //Returns true if there is a div in the document.
div.HasChild("p") //Returns true if there are p tags in the div's children.
div.HasChild("p", true) //Returns true if there are p tags in the div's children, recursive.
div.HasAttr("attr") //Returns true if the div has the attr attribute.
div.HasAttrValue("attr", "value") //Returns true if the div has the attr attribute set to "value".
```

And methods to remove attributes :

```
div.RmClass("class") //Removes the div's class attribute.
div.RmAttr("attr") //Removes the div's "attr" attribute.
```

Once your HTML tags are ready, you can add them to the body container this way :

```
body.Put(div) //Adds the div to the body 
body.Make(div, div2, div3) //Adds a list of divs to the body
```

And then we add our HTML to the document :

```
doc.Put(body) //Adds the body container to the document. 
doc.Make(header, body, footer) //Adds several containers to the document
```

Now the HTML is ready. Let's take a look at css styling with our
sheet object :

```
style := *frongo.NewStyle("selector", "modifiers", "values") //Creates a style object
sheet.Put(style) //Adds a style object to the sheet
sheet.Make(style, otherStyle) //Adds many style objects to the sheet
sheet.Style("selector", "modifiers", "values") //Adds a new style, this method can be chained
```

Then we add the sheet to the document :

```
doc.Style(sheet, otherEventualSheet) //Adds a list of sheets to the document
```

Now, the document is ready. 
You can either get a HTML template out of it or save it as a file :

```
doc.Render() //Returns the complete HTML template
body.Render() //Returns the body's HTML template
sheet.Render() //Returns the sheet's HTML template

frongo.Create("filename", doc.Render()) //Creates a file for the document, indirect way

doc.Create("filename") //Creates a file for the document
body.Create("filename") //Creates a file for the body
sheet.Create("filename") //Creates a file for the sheet
```

Now everything is set and you can use either the resulting file or the resulting template.
And everything was made using only Go !