Base64 Encoding

Go provides built-in support for base64 encoding/decoding.
	


This syntax imports the encoding/base64 package with the b64 name instead of the default base64. It’ll save us some space below.
	


Here’s the string we’ll encode/decode.
	


Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder. The encoder requires a []byte so we convert our string to that type.
	


Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	



This encodes/decodes using a URL-compatible base64 format.
	


The string encodes to slightly different values with the standard and URL base64 encoders (trailing + vs -) but they both decode to the original string as desired.
	



