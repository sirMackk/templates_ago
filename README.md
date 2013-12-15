A very simple package that helps with golang templates. It simply takes the name of the controller/function and looks for a template with the same name and uses that.

Example usage:

    var templates = templates_ago.NewTemplates()
    templates_ago.LoadTemplates("templates/", templates)

    // in controller 
    err = templates.Execute(w, struct_or_map)


Copyright (C) 2013 sirMackk

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.
