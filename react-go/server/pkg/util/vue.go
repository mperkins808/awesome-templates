package util

import (
	"io/fs"
	"net/http"
	"regexp"
	"text/template"

	vueglue "github.com/mperkins808/vite-go"
	log "github.com/sirupsen/logrus"
)

func serveOneFile(w http.ResponseWriter, r *http.Request, vueData *vueglue.VueGlue, uri, contentType string) {
	strippedURI := uri[1:]
	buf, err := fs.ReadFile(vueData.DistFS, strippedURI)
	if err != nil {

		append := "public/" + strippedURI
		if vueData.Environment == "production" {
			append = "dist/" + strippedURI
		}

		buf, err = fs.ReadFile(vueData.DistFS, append)
	}

	// If we ended up nil, render the file out.
	if err == nil {
		// not an error; letting the error case fall through
		w.Header().Add("Content-Type", contentType)
		w.Write(buf)
		return
	}
	log.Error(err)

	// Otherwise, we cannot handle it, so 404 it is.
	w.WriteHeader(http.StatusNotFound)
}

func PageWithAVue(w http.ResponseWriter, r *http.Request, vueData *vueglue.VueGlue) {
	re := regexp.MustCompile(`^/([^.]+)\.(svg|ico|jpg|png)$`)
	matches := re.FindStringSubmatch(r.RequestURI)
	if matches != nil {
		if vueData.Environment == "development" {
			log.Printf("redirecting request to vite dev server")
			url := vueData.BaseURL + r.RequestURI
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
			return
		} else {
			// production; we need to render this ourselves.
			var contentType string
			switch matches[2] {
			case "svg":
				contentType = "image/svg+xml"
			case "ico":
				contentType = "image/x-icon"
			case "jpg":
				contentType = "image/jpeg"
			case "png":
				contentType = "image/png"
			}

			serveOneFile(w, r, vueData, r.RequestURI, contentType)
			return
		}

	}

	// our go page, which will host our javascript.
	t, err := template.ParseFiles("./template/template.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, vueData)
}
