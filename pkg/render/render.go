package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//functions is not built into the templates_functions
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	/*	_, err := RenderTemplateTest(w)
		if err != nil {
			fmt.Println("error getting template cache:", err)
		}*/

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error wringing template to browser", err)
	}

	/*	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template:", err)
			return
		}*/
}

//CreateTemplateCache creates a cache of available templates from the templates directory
func CreateTemplateCache() (map[string]*template.Template, error) {

	myTemplateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myTemplateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//log.Println("page is currently",page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myTemplateCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myTemplateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myTemplateCache, err
			}
		}

		myTemplateCache[name] = ts
	}
	return myTemplateCache, nil
}
