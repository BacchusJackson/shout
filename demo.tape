Set Theme "Catppuccin Mocha"

Set FontSize 24
Set FontFamily "JetBrainsMonoNL Nerd Font Mono"
Set Margin 20
Set MarginFill "#89B4FA"
Set BorderRadius 10

Set Width 1500

Output bin/shout-demo.gif

Type "shout version"
Sleep 1
Enter
Sleep 1

Type "shout log some informational log message"
Sleep 1
Enter
Sleep 1

Type "shout warn some warning message"
Sleep 1
Enter
Sleep 1

Type "shout error some error message"
Sleep 1
Enter
Sleep 1

Type "shout alert http://localhost:3000/webhook 'name=example alert' 'message=this is a custom field for this example webhook' "
Sleep 1
Enter
Sleep 5

