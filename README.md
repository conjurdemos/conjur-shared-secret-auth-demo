# Conjur Rotation Demo

## Prerequisites

1. Docker and Docker Compose

## Running the Demo

1. Start Conjur
    This will create the containers for Conjur and the application services and start the
    Conjur service.
    ```
    $ ./start
    ```

2. Load the Conjur policy
    ```
    $ ./load-policy
    ```

3. Start the **helloworld** application service
    ```
    $ ./start-helloworld
    ```

4. Call the **helloworld** service, using Conjur to deliver the **helloworld** application secret
    ```
    $ ./connect
    ```

5. Rotate the **helloworld** application secret and restart the service.
    > This will cause **helloworld** to update Conjur with the new secret
    ```
    $ ./rotate

    # Executing connect again will use the updated app secret from Conjur
    $ ./connect
    ```

6. An unauthorized service request can be issued using connect with `--fail`
    ```
    $ ./connect --fail
    ```
