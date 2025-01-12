//Explanation:
//PaymentService Interface: Defines a common interface for all payment service implementations, allowing for loose coupling.
//Concrete Implementations (PaypalService, StripeService): Each implements the PaymentService interface, providing specific payment processing logic.
//PaymentProcessor: Acts as the intermediary layer that decouples the client from the payment services. It uses the PaymentService interface to process payments, allowing easy switching between different payment services.
//Client Code: Demonstrates how to use PaymentProcessor to handle payments through different services by selecting the service at runtime, showcasing flexibility and modularity.
