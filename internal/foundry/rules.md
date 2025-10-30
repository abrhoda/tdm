# Rules
This document stands as a breakdown of the clusterfuck that is the rules in `foundryvtt/pf2e`. These rules heavy leverage that lists in js can have mixed types. To handle this, there will need to be a lot of `maybeIntAsString` types to handle places where the list can have both strings and ints, for example.

## Break down.

### Predicate
`predicate` is always a ton level key in a rule. It is a list that can contain:
1. empty
2. strings
3. predicateComplexObj

Where a `predicateComplexObj` is a struct that has only 1 key of the following:
1. Or
2. And
3. Not
4. Gte
5. Lte
6. Gt
7. Lt
8. Eq

### PredicateComplexObjs
Below are the per down, per key name, of the sub objects for the predicateComplexObj

`Or` can have list of:
1. strings
2. predicateComplexObj

`And` can have list of:
1. strings
2. predicateComplexObj

`Not` can have list of:
1. string
2. predicateComplexObj

`Gte` can have list of:
1. strings
2. ints

`Lte` can have list of:
1. strings
2. ints

`Gt` can have list of:
1. strings
2. ints

`Lt` can have list of:
1. strings
2. ints

`Eq` can have list of:
1. strings
