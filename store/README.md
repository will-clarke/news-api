# Store

The store satisfies the `ServerInterface`, which is something our openapi client
automatically generates.

It's a very stupid implementation of the interface and only stores the data
in-memory. In reality we'd want something a bit more persistent and durable.

Ideally it would be cool if we could split this interface into two (and have an
`ArticleStore` and `FeedStore`), but I think openapi prefers a single
file. ¯\_(ツ)_/¯
