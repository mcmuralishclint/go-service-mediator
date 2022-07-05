# GO-SERVICE-MEDIATOR
## UNDER CONSTRUCTION

This is a service orchestration package based on Golang and the motive of this project is to allow anyone to define their project's API contracts within this service and make calls between microservices via this package.

**Advantages**
- A centralized place for storing your API contracts
- Easy to test your microservcies (unit test, integration test)
- Easy to do API version migration. As all the API contracts are defined in one place, donstream services can make use of feature flags to easily switch traffic between different API versions in clean manner.
- Swagger documentation process will be smoother. In house or external resources can easily comprehend the intention of the APIs by just reading the test cases.
- Changes in API contracts can be easily identified.

**Techniques Used**
- Using GRPC for network calls: Proven to be way faster than REST APIs. Since this is an orchestration service, in any way, forwarding traffic through this service should not impact performance.
- The configs are stored in a yaml file in a particular format and users are required to define their yaml file after cloning this project.
