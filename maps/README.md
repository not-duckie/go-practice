# Maps

## Restrictions

1. The key types need to be test for equality thus they can be slices,maps or other functions 


## Number of ways to declare a map

1. Creating map without giving any initial value

```
hashmap := make(map[string]string)
```

2. Creating an empty map

```
hashmap := map[string]int{}
```

3. Creating map with initial values

```
hashmap := map[string]int{
	"Calofornia": 69,
	"sex i had": 0,
		}
```


## IMPORTANT NOTICE

the order of a hashmap can change after adding a new key to it. So do not rely on it.
