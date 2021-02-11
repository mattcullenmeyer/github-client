## GitHub Stargazer Client

The GitHub Stargazer Client is a Command Line Interface (CLI) that interacts with the GitHub Stargazer Server, which returns how man stars a given GitHub repository has received.  

### Getting Started

Get started by cloning the repo:  
`$ git clone https://github.com/mattcullenmeyer/github-stargazer-client.git`  

Navigate to the root directory of the project and enter the following command:  
`$ go mod init github.com/mattcullenmeyer/github-stargazer-client`

### Running the Client

Once cloned, navigate to the root directory of the project. The CLI is executed by entering `go run main.go` followed by an indefinite number of `<organization>/<repository>` string arguments.   

For example, to return the number of stars that `https://github.com/mattcullenmeyer/github-stargazer-client` has received, enter:  
`$ go run main.go mattcullenmeyer/github-stargazer-client`  

If you want to see the number of stars that multiple repos have received, try something like this:  
`$ go run main.go mattcullenmeyer/github-stargazer-client mattcullenmeyer/github-stargazer-server` 

### Running Tests

The GitHub Stargazer Client includes unit tests that you can run to ensure it's working properly.  

Run the tests by entering the following command in the root directory of the project.  
`$ go test`