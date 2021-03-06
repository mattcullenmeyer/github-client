## GitHub Stargazer Client

The GitHub Stargazer Client is a Command Line Interface (CLI) that interacts with the GitHub Stargazer Server, which returns how man stars a given GitHub repository has received.  

### Getting Started

Get started by cloning the repo:  
`$ git clone https://github.com/mattcullenmeyer/github-stargazer-client.git`  

Once cloned, navigate to the root directory of the project:   
`$ cd github-stargazer-client`  

### Running the Client

The CLI is executed by entering `go run main.go` followed by an optional url flag and an indefinite number of `<organization>/<repository>` string arguments.   

If you are running the Server from Docker, which has a specified url of `http://localhost:8080`, you can get the number of stars for `https://github.com/mattcullenmeyer/github-stargazer-client` by entering:  
`$ go run main.go mattcullenmeyer/github-stargazer-client`  

Because `http://localhost:8080` is the default url, the following command will yield the same result :  
`$ go run main.go -url=http://localhost:8080 mattcullenmeyer/github-stargazer-client`  

You will however need to specify a url path if you are running the Server from Kubernetes since the url is unknown. For example:  
`$ go run main.go -url=http://192.168.49.2:30893 mattcullenmeyer/github-stargazer-client`  

If you want to see the number of stars that multiple repos have received, try something like this:  
`$ go run main.go -url=http://192.168.49.2:30893 mattcullenmeyer/github-stargazer-client mattcullenmeyer/github-stargazer-server` 

### Running Tests

The GitHub Stargazer Client includes unit tests that you can run to ensure it's working properly. Make sure the GitHub Stargazer Server is first running in **Docker** before running the Client tests.   

Run the tests by entering the following command in the root directory of the project.  
`$ go test`