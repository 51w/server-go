# server-go


忽略警告（不推荐）：在Git命令行中运行以下命令以忽略警告：
git config --global core.safecrlf false

转换行尾符为LF：如果你在Windows上开发，可以考虑将行尾符转换为LF。在Git命令行中运行以下命令：
git config --global core.autocrlf input

转换行尾符为CRLF：如果你在Unix/Linux上开发，并且希望行尾符被转换为CRLF，可以运行以下命令：
git config --global core.autocrlf true
