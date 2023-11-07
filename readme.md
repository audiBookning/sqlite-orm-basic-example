# sqlite orm basic example


Just a basic example of how to do a custom little sqlite orm for use in simple and limited use cases.
## Notes

The "Table struct" is

``` go
type Book struct {
	ID           int
	BookName     string
	CleanedTitle string
	URL          string
}
```

with also 2 date fields for added and edited record.

The "orm" package is called books and is directly dependant on the Book struct, which means that the code will have to be diredctly changed for any other data structure.

Why? It seems that addinmg more abstrction to it is one of those rabbit holes and then i am not sure why one won't use a normal ORM package if one must enter in so many complexities.


## Disclaimer

There is no intention in turning this in any sort of maintained project. Anyone is welcome to copy or fork it to do whatever they want.
