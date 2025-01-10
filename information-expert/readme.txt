/Explanation:
//Product Class: Represents a product with its name and price. It encapsulates data relevant to the product.
//DiscountCalculator Class: An information expert dedicated to calculating discounts. It has the necessary information (product price and discount percentage) to perform its calculations.
//Checkout Class: Manages a collection of products and calculates the total cost. It also has the logic to display the checkout summary, demonstrating that it has responsibility over products and their total cost.
//Client Code: In the main function, products are added to the checkout, and a summary is printed, showing how the responsibility for calculating the discount is appropriately assigned to the DiscountCalculator.

Information Expert
Definition: Assign responsibility to the class that has the necessary information to fulfill it.
Purpose: Promotes low coupling and high cohesion by ensuring that methods are allocated to classes that have the necessary data.
Example: A Customer class should handle tasks related to customer data, such as validating a customer’s credit because it owns that data.