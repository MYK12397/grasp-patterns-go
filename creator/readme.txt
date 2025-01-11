//Explanation:
//Item Struct: Represents individual items in a store.
//NewItem Function: Serves as a factory function, but its responsibility could reside in Cart to closely tie the creation of an item to its usage context.
//Cart Struct: Acts as a Creator, responsible for creating Item instances through the AddItemToCart method and managing a collection of items.
//Client Code: Demonstrated in the main function, where items are added directly to the cart, showcasing how responsibility for item creation is contained within the Cart class.
//Summary:
//The Creator pattern helps in putting responsibility for the creation of related objects within a class that aggregates or uses them, thereby guiding object creation in a way that enhances cohesion and reduces unnecessary dependencies. By following this principle, you can create maintainable and scalable systems due to well-defined class responsibilities.


Creator Pattern
The Creator pattern is another key GRASP (General Responsibility Assignment Software Principles) principle that guides the assignment of responsibility for creating instances of certain classes. This pattern suggests that responsibility for creating an object should be assigned to a class that has a meaningful relationship with the object to be created.

Key Features:
Relational Logic: The class that creates an object has a strong relationship with that object, such as aggregation, composition, or usage.
Simplification: Reduces the complexity of object creation logic by centralizing it within a relevant class.
Cohesion: Increases cohesion by ensuring that related responsibilities are grouped together.
Benefits:
Reduced Coupling: Assigning creation responsibility strategically avoids unnecessary dependencies between classes.
Increased Reusability: By isolating creation logic, it fosters code reuse and more efficient object lifecycle management.
Clear Responsibility: Clarity is achieved as classes directly related to an object’s lifecycle handle its creation.
Potential Problems:
Complex Logic in Creator: If overused, might lead to bloated creator classes that do too much.
Misassigned Creation Logic: Misidentifying which class should be the creator can lead to poor design and lower cohesion.
