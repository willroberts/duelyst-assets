# Process

Documents the process of mapping and using Duelyst's internal APIs, with the
goal of retrieving and displaying my personal card collection.

## Phase 1: How can I see the game's source code?

Duelyst is written in Javascript, which means that it's probably not compiled.
I installed Duelyst via Steam, so I found the game's source code here:

```
C:\Program Files\Steam\steamapps\common\duelyst\resources\src\duelyst.js
```

I ran `js-beautify` on this file, and stored it in this repo at
`src/duelyst.js`. You can get the `js-beautify` program from PyPI:

```
pip install jsbeautifier
```

Now we have the source code in a readable format!

## Phase 2: What are the relevant routes for the API?

After un-minifying the source code, I used the `findurls` program to scrape all
unique URLs from the source code, filtering only the routes for inventory:

```
go run cmd/findurls/main.go | grep inventory
```

These looked particularly relevant:

```
https://duelyst-production.firebaseio.com/user-inventory/
https://play.duelyst.com/api/me/inventory/card_collection
https://play.duelyst.com/api/me/inventory/card_collection/duplicates
https://play.duelyst.com/api/me/inventory/card_collection/read_all
https://play.duelyst.com/api/me/inventory/card_collection/soft_wipe  # scary :O
```

Duelyst uses Firebase, a managed PubSub service, presumably for sending
messages to/from their servers and possibly between clients. I'm not sure how
this fits in yet, so we'll focus on the routes marked `/api/`.

Looking at the surrounding source code, most of these routes are used in `POST`
or `DELETE` requests for purposes other than retrieving a collection. There may
not be an API endpoint for retrieving the collection.

Perhaps the collection data is sent between Duelyst's backend servers and my
client via Firebase? To be continued...

### Side Quest: How do I authenticate to the API?

Sending a request to the `/card_collection` route fails:

```
curl -s https://play.duelyst.com/api/me/inventory/card_collection
```

The response is in HTML format instead of JSON (ouch!), but the title gives us
some useful information:

```
<title>No authorization token was found</title>
```

Searching for `token` in the source code yields this bit of information:

```
return localStorage.getItem("token")
```

It looks like the token is stored in a browser's local storage. I launched the
[Duelyst web client](http://play.duelyst.com) in a browser, logged in, and then
looked at local storage in the dev console. The `token` key contained my token,
so I copied it locally.

The `firebase:session:duelyst-production` key contained some metadata alongside
the token, including an expiration timestamp. My token was valid for around two
weeks.

Well... how do I use the token? Not the `Authorization` header.
