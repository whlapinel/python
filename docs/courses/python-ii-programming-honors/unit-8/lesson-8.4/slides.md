---
marp: true
theme: default
class: lead
paginate: true
---

{%- comment -%} Disable Liquid processing {%- endcomment -%}

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Agenda

- No exam this week
- Catch-up and/or work ahead
  - Please check your grades in Canvas for anything missing
  - Lesson 8.4 is posted for working ahead

# Flask Templates and Components

# Introduction to Flask Templates

- Flask uses **Jinja2** templating engine.
- Allows you to embed Python code in HTML.
- Helps create dynamic web pages with ease.

# Setting Up Templates

1. Create a `templates/` folder in your Flask project directory.
2. Add HTML files in the `templates/` folder.

Example folder structure:

```plaintext
project/
  app.py
  templates/
    index.html
    about.html
```

# Rendering a Template

- Use the `render_template` function to display templates.

**Example:**

```python
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html', title="Home")

if __name__ == '__main__':
    app.run(debug=True)
```

---

# Using Template Variables

- Pass variables to templates via `render_template`.

**app.py:**

```python
@app.route('/about')
def about():
    return render_template('about.html', title="About", name="John")
```

# **about.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    <h1>Welcome, {{ name }}!</h1>
</body>
</html>
```

# Template Components (Reusable HTML)

- Use `{% include %}` to insert reusable components.

# Step 1: Create the Component

**templates/header.html:**

```html
<header>
    <h1>My Website</h1>
    <nav>
        <a href="/">Home</a> |
        <a href="/about">About</a>
    </nav>
</header>
```

# Step 2: Include the Component in Pages

# **templates/index.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    {% include 'header.html' %}
    <h2>Welcome to the Homepage!</h2>
</body>
</html>
```

# **templates/about.html:**

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    {% include 'header.html' %}
    <h2>About Us</h2>
</body>
</html>
```

# Benefits of Components

- **Reusability**: Avoid duplicating HTML code.
- **Consistency**: Ensure uniform styling across pages.
- **Maintainability**: Update the component once, reflect changes everywhere.

# Next Steps

1. Practice creating templates and components.
2. Explore other Jinja2 features:
   - Conditionals (`{% if %}`)
   - Loops (`{% for %}`)
3. Build a simple multi-page Flask application.
