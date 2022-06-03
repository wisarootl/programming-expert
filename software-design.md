# Divide And Conquer

**Divide And Conquer** : divide into sub-problem / sub-system

# Cohesion

**Cohesion** : degree of similar thing together (the higher, the better)

Degree to which things grouped together are similar. Usually, when programming, you would like to increase cohesion. There are many different types of cohesion:

- Functional Cohesion : with no side effects. (side effect = thing outside the function)
- Layer Cohesion : high level layer access low level layer.
- Communicational : facilities for operating on the same data are kept together. e.g., class
- Sequential Cohesion : a set of procedures which work in sequence (Data-preprocessing P/L)
- Procedural Cohesion : similar to Sequential Cohesion but do not have to work with the same data. It just have to work in order.
- Temporal Cohesion : procedures used in the same general phase of execution like initialization and cleanup are kept together. (Open and Close the file)
- Utility : related utilities are kept together (when there is no way to group using strong cohesion). For example, putting helper function into utils.py

# Coupling

**Coupling** : degree of interdependence between things (the lower, the better)

degree of interdependence between things. In programming, it is typically desirable to reduce coupling. There are many different types of coupling:

- Content : a component secretly modifies the internal data of another component.
- Stamp : too specific parameter type in function / method
- Routine Call : function / method call other functions or methods
- Type-use : use of globally defined data types.
- Inclusion/Import : importing unnecessary components.
- External : a dependency on something outside of the scope of the system like an operating system, shared library, hardware etc.

# Abstraction

**Abstraction** : increasing generality

For example: for loop in Python can work with any iterable object.

# Reusability

**Reusability** : increasing reusability

# Reuse

**Reuse** : not reinvent the wheel. use existing framework rather than create your own

# Flexibility

**Flexibility** :

- reduce coupling and increases cohesion. (separation between frontend and backend)
- increase abstraction
- not hardcode but use variables
- make reusable code or reuse available code

# Obsolescence

**Obsolescence** :

- avoid early version of software
- avoid software with environment specific (only work on windows / linux)
- avoid un-document features / not traditional used / will de depreciated in the future
- hardware / software with no support (old hardware / software)

# Portablity

**Portablity** : run on many platform/environment as possible

# Testability

**Testability** : automated test / write the code that can be easier to test (e.g., many small functions are better than one large function)

# Defensibility

**Defensibility** : preventing misuse of our code

- precondition : check type of argument
