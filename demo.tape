# VHS demo script for Namigo CLI
#
# This script demonstrates the Namigo CLI tool by running various search commands.
# It showcases package, DNS, and email search functionalities.

Output images/demo.gif

Require namigo

Set Shell "zsh"
Set FontSize 18
Set Width 1280
Set Height 720
Set TypingSpeed 100ms

Type "namigo generate prompt" Enter
Sleep 500ms Type "a" Enter
Sleep 500ms Type "b" Enter
Sleep 500ms Type "c" Enter
Sleep 500ms Type "d" Enter
Sleep 1s

Type "namigo search package hello" Enter
Sleep 4s

Type "namigo search dns hello" Enter
Sleep 4s

Type "namigo search email hello" Enter
Sleep 4s

Type "exit" Enter
Sleep 0.5s
