A very simple package that helps with golang templates. It simply takes the name of the controller/function and looks for a template with the same name and uses that.

Example usage:

    var templates = templates_ago.NewTemplates()
    templates_ago.LoadTemplates("templates/", templates)

    // in controller 
    err = templates.Execute(w, struct_or_map)

