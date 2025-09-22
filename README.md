# Rum
A streamlined http router and context handler for Golang Microservices.

This will largely be based on the Gin routing system.

### Why?
I was importing Gin throughout my services as a dependancy and realized I only use a small slice
of what it offers. 

## What I need out of an HTTP Router?
1. Ability to configure routes and route groups.
2. Ability to deploy middleware.
3. UUID Generation inside Context for route tracing and logging.
4. JSON Writers

## Side Notes
While going through this I realized that one huge advantage of Gin's system is that they
use the bytedance/json library for JIT compiled JSON Marshalling and UnMarshalling.
The benefits of this are substantial when dealing with large JSON structures..so I may want to include 
optional support for it.
For the sake of simplicity I'm keep as is for now but will return to this.
