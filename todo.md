- Context
- Type assertion & switches
- Middleware
x Panic & Recovery
x new  => Initialize struct with zero values.
x Loop tags
x goto 


Re: Context: It's a tree structure, where you can spawn multiple children off of each parent to hand over to other functions and services, and if the parent is canceled, all descending children are also canceled. The context also serves as a key-value store, which can be used to carry info forward.