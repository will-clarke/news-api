# Importer

- This package's responsibility is to receive a new feed and then:
1. parse the feed
2. if valid, store the feed in the db
3. for each article in the feed, store the article

- I should probably actually test this. I'd use an interface that can
  `ParseURL`and mock that out somehow.
