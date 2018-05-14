# Conjur Enterprise Shared Secret Authentication Demo
This example demonstrates using Conjur Enterprise 
to store, rotate, and fetch an application secret required to 
access a protected system application. 

This example is the same as the Conjur Open Source rotation demo,
adapted to work with the Conjur Enterprise Appliance.

## Prerequisites

To run this system demo, you will need:

1. Docker and Docker Compose

2. Access to a Conjur Enterprise Appliance Docker container image

## Running the Demo

1. Navigate to the `enterprise/` directory in the depo repository

    ```bash
    $ cd enterprise
    ```

2. Specify the Conjur Enterprise Application container image to use as
    the environment variable `CONJUR_APPLIANCE_IMAGE`

    ```bash
    $ export CONJUR_APPLIANCE_IMAGE=my-registry.net/conjur-appliance:4.9-stable
    ```    

3. Start Conjur Enterprise
    This will create the containers for Conjur and the application services
    and start the Conjur service.

    You will be prompted to verify the Conjur Enterprise certificate and log
    into the appliance to load the policy.

    > NOTE: This command will also clean up any previous application state and
    > reset the system

    ```bash
    $ bin/start-conjur
    ```

2. Load the Conjur policy

    > NOTE: You can review the policy [here](cli/policy.yaml).

    ```bash
    $ bin/load-policy
    ```

3. Start the **helloworld** application service

    ```bash
    $ bin/start-helloworld
    ```

4. Call the **helloworld** service, using Conjur to deliver the **helloworld** application secret

    ```bash
    $ bin/run-consumer
    ```

5. Restart the **helloworld** service to rotate its secret.

    > NOTE: This will cause **helloworld** to update Conjur with the new secret.
    > You should expect to see `bin/run-consumer` display a different API secret
    > from the previous execution.

    ```bash
    $ bin/start-helloworld

    # Executing connect again will use the updated app secret from Conjur
    $ bin/run-consumer
    ```

6. An unauthorized service request can be sent by running the consumer
   with `--fail`

    ```bash
    $ bin/run-consumer --fail
    ```
