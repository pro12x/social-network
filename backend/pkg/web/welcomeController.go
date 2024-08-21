package web

import (
	"backend/pkg/globale"
	"backend/pkg/utils"
	"html/template"
	"net/http"
)

type Data struct {
	Version  string
	Author   string
	Email    string
	Github   string
	Linkedin string
	Endpoint []Endpoint
}

type Endpoint struct {
	Name string
	Urls []string
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/swagger" {
		http.NotFound(w, r)
		utils.LoggerError.Println(utils.Error + "404 Not Found" + utils.Reset)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	page := `
	<!DOCTYPE html>
	<html lang="en" data-bs-theme="dark">
		<head>
  			<title>My Swagger</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
			<style>
				a {
					text-decoration: none !important;
				}
			</style>
		</head>
		<body>

			<nav class="navbar navbar-expand-sm bg-dark navbar-dark container d-flex justify-content-center">
				<ul class="navbar-nav me-auto">
					<li class="nav-item">
						<a class="nav-link active fw-bold" href="javascript:void(0)">Social Network</a>
					</li>
				</ul>
				<ul class="navbar-nav">
					<li class="nav-item">
						<a class="nav-link" href="javascript:void(0)">Login</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="javascript:void(0)">Register</a>
					</li>
				</ul>
			</nav>
			
			<div class="container mb-5">
				<h1 class="mt-4 p-2 bg-primary text-white rounded">About</h1>
				<ul class="list-group list-group-flush">
					<li class="list-group-item"><b>Author:</b>&nbsp;<a class="btn btn-link badge bg-secondary" href="https://www.google.com/search?q=%22franchis+janel+mokomba%22" target="_blanc">{{.Author}}</a></li>
					<li class="list-group-item"><b>Email:</b>&nbsp;<a class="btn btn-link badge bg-secondary" href="mailto:janelaffranchis@gmail.com">{{.Email}}</a></li>
					<li class="list-group-item"><b>Github:</b>&nbsp;<a class="btn btn-link badge bg-secondary" href="https://github.com/pro12x" target="_blank">{{.Github}}</a></li>
					<li class="list-group-item"><b>Linkedin:</b>&nbsp;<a class="btn btn-link badge bg-secondary" href="https://www.linkedin.com/in/franchisjanelmokomba" target="_blank">{{.Linkedin}}</a></li>
					<li class="list-group-item"><b>Version:</b>&nbsp;<a class="btn btn-link badge bg-secondary">{{.Version}}</a></li>
				</ul>

				<h1 class="mt-4 p-2 bg-primary text-white rounded">Endpoints</h1>
                <div id="accordion">
					{{range .Endpoint}}
                	<div class="card mb-1">
                    	<div class="card-header">
                        	<a class="btn" data-bs-toggle="collapse" href="#{{.Name}}">
                            	{{.Name}}
                            </a>
                        </div>
                        <div id="{{.Name}}" class="collapse" data-bs-parent="#accordion">
                        	<div class="card-body">
								<ul class="list-group list-group-flush">
									{{range .Urls}}
									<li class="list-group-item fw-bold list-group-item-info rounded">{{.}}</li>
									{{end}}
								</ul>
                          	</div>
                      	</div>
                    </div>
					{{end}}
                </div>
			</div>

			<div class="mt-2 p-2 text-white text-center">
				<p>&copy; 2024 by Franchis Janel MOKOMBA. All Rights Reserved. Social Network is Powered by Zone01 Dakar</p>
			</div>

		</body>
	</html>
	`

	data := Data{
		Version:  "1.0.0",
		Author:   "Franchis Janel MOKOMBA",
		Email:    "janelaffranchis@gmail.com",
		Github:   "@pro12x",
		Linkedin: "@franchisjanelmokomba",
		Endpoint: convert(globale.Endpoints),
	}

	t, err := template.New("webpage").Parse(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		utils.LoggerError.Println(utils.Error + err.Error() + utils.Reset)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		utils.LoggerError.Println(utils.Error + err.Error() + utils.Reset)
		return
	}
}

func convert(endpoints map[string][]interface{}) []Endpoint {
	var endpointsList []Endpoint
	for k, v := range endpoints {
		var urls []string
		for _, u := range v {
			urls = append(urls, u.(string))
		}
		endpointsList = append(endpointsList, Endpoint{Name: k, Urls: urls})
	}
	return endpointsList
}
