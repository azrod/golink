package short

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var tmpl404 = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>404 Not Found</title>
		<style>
			html,
			body {
				height: 100%;
				margin: 0;
				padding: 0;
			}
			
			.container {
				height: 100%;
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
			}
			
			.title {
				font-size: 3rem;
				font-family: sans-serif;
				color: #1a202c;
			}
			
			.subtitle {
				font-size: 1.5rem;
				font-family: sans-serif;
				color: #718096;
			}

			.redirect {
				font-size: 1rem;
				font-family: sans-serif;
				color: #718096;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="title">Link Not Found</div>
			<div class="subtitle">Sorry, the link you are looking for does not exist.</div>
			<div class="redirect">Redirecting you to the home page in 5 seconds... or <a href="/u/">click here</a>.</div>
			<script>
				setTimeout(function() {
					window.location.href = "/u/";
				}, 5000);
			</script>
		</div>
	</body>
	</html>
	`

func (s *Short) handleHTML404(c echo.Context) error {
	return c.HTMLBlob(http.StatusNotFound, []byte(tmpl404))
}
