SHA256 Hashes

SHA256 hashes are frequently used to compute short identities for binary or text blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate’s signature. Here’s how to compute SHA256 hashes in Go.
	


Go implements several hash functions in various crypto/* packages.
	


	


Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	

This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	



Running the program computes the hash and prints it in a human-readable hex format.
	


You can compute other hashes using a similar pattern to the one shown above. For example, to compute SHA512 hashes import crypto/sha512 and use sha512.New().
	

Note that if you need cryptographically secure hashes, you should carefully research hash strength!