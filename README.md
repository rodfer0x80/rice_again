# rice_again
## restful api
````
# methods
## create
GET
POST
## read
GET
POST
## update
GET
POST
PUT
## delete
GET
POST
DELETE
````
````
## create
/$name
/$name/new
/$name/create
## read
/$name
## update
/$name
/$name/edit
/$name/update
## delete
/$name
/$name/delete
/$name/del
````
## graphql
````
# methods
GET
POST
````
````
# queries
/gql?q=$name
````

## architecture
````
separate functions by methods
create separate payloadlists for each method as sets so there are no duplicates
there wordlists are made by adding each iteration of path from default/custom wordlist for resful apis
this is for restful ofc, graphql is much easier just needs a different wordlist and can use a single payloadlist
get status code from requests
to build api map
````
