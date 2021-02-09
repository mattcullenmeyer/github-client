## GitHub Stargazer Client

The GitHub Stargazer Client is a Command Line Interface (CLI) that interacts with the GitHub Stargazer Server, which returns how man stars a given GitHub repository has received.  

Get started by cloning the repo:  
`$ git clone https://github.com/mattcullenmeyer/github-client.git`  

Once cloned, navigate to the root directory of the project. The CLI is executed by entering `go run main.go` followed by an indefinite number of `<organization>/<repository>` string arguments.   

For example, to return the number of stars that `https://github.com/mattcullenmeyer/github-client` has received, enter:  
`$ go run main.go mattcullenmeyer/github-client`  

If you want to see the number of stars that multiple repos have received, try something like this:  
`$ go run main.go mattcullenmeyer/github-client mattcullenmeyer/anaplan mattcullenmeyer/tinytrader` 