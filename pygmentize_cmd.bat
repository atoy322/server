pygmentize -l go -f html -O linenos=True,style=dracula -o main.go.html main.go
pygmentize -S dracula -f html -a .highlight > html_dracula.css
