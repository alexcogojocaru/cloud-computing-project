# cloud-computing-project

## driver service
- create a driver account -> request to auth service
- on every login check if the credentials are valid -> request to auth service
- run a server on driver's device (websocket, grpc)
    - update the driver's location every `POLLING_TIME` seconds
    - listen for incoming requests from riders
        - can accept/decline the ride
        - request the route from google maps api (needs start & end destinations)
        - when the ride starts, a connection between the driver and the rider is created
        - trunc the ETA between `[LOW, UPPER]` for simulation purposes

## rider service
- create a rider account -> request to auth service
- on every login check if the credentials are valid -> request to auth service
- the user's location is automatically extracted by the service
- the user needs to select the destination and the payload will be sent to the ride service that returns the closest driver and creates the connection between them

## auth service
- manages/creates accounts

## ride service
- receives updates with a driver's location
- on a rider's request, it searches for the closest driver and creates a connection between them
