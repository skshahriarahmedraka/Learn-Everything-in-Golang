XML

Go offers built-in support for XML and XML-like formats with the encoding.xml package.
	

Plant will be mapped to XML. Similarly to the JSON examples, field tags contain directives for the encoder and decoder. Here we use some special features of the XML package: the XMLName field name dictates the name of the XML element representing this struct; id,attr means that the Id field is an XML attribute rather than a nested element.
	



Emit XML representing our plant; using MarshalIndent to produce a more human-readable output.
	

 
To add a generic XML header to the output, append it explicitly.
	

Use Unmarshal to parse a stream of bytes with XML into a data structure. If the XML is malformed or cannot be mapped onto Plant, a descriptive error will be returned.
	



The parent>child>plant field tag tells the encoder to nest all plants under <parent><child>...