# Decisions

* For simplicity reasons
    * I haven't implemented a `sign in` endpoint.
    * I chose to send the `user` within the payload. In a real case, I would have used jwt. Because of this decision, the `resume` endpoint is a `POST` instead of a `GET`.
    * I chose `redis` as the main storage. In a real minesweeper I would have chosen another `NoSQL` data base, particularly `MongoDB`.

* Known bugs
    * When creating a game, in the response you can find the `end` attribute with the default value for `time`. The field should be omitted.
    * `SwaggerUI` can't access the `aws` instance. That's why I also give you some `curl`s. 

* Things I consider important
    * I wrote a custom `GameResponseDto` because I think that the user should never *see* if a cell is mined or not. With this `dto` I can hide that property.
    * Even though a `cell` shouldn't know anything about its adjacent cells, I added the attribute `NearMines` because in another way, the `board` should keep a `map` with this information. So I preferred to give it to the `cell`.  
    * By design, a `flagged cell` can't be explored.
    * Internal errors are never specified to the user.
    * The `flag` endpoint is a `PUT` because it is idempotent.

* Extras
    * I implemented a memory cache.
    * At the beginning of the project, I wondered if using `DDD` with `Go` would be nice. I wasn't sure, even though `DDD` is language agnostic. So I used `DDD` as a proof of concept. The results were nice in my opinion, I will give it another try in the future.
