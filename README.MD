# CARS VIEWER

Cars Viewer is a project that includes a web server and a web interface working together with an API.
The web includes a gallery of cars, a search bar and a filter menu. It also allows you to compare cars, create a favourite list and review the last comparison made by the user.




## Usage

To run the application, follow these steps:
- Clone the Repository.
- Install the API. This API provides the data for the car models and needs to be installed separately. Follow the instructions in the API's README file to install it. (cars/api/readme)
- In terminal, navigate to the API directory (cars/api) and start the API using the following command: `make run`
- In another terminal(split terminal), navigate to the root directory for the project (/cars) and start the server by running: `go run .`
    - by default server is running on port 8080, but you can use -p flag to specify the port that the server should run on: `go run . -p=8080` 
- Finally, access your browser and go to: http://localhost:port to get in the website.


## Endpoints

Server exposes following endpoints:

- `/` Main gallery of all cars cards
- `/car` Info on the specific car. Should be called with url query, otherwise shows the error
- `/filter` gets cars from api based on filter parameters
- `/compare` shows info on multiple cars, previously chosen by the user
- `/last` shows cards for up to 5 last visited cars
  