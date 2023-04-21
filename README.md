![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# 8menu

## what?

8menu is a simple menu utility inspired by the behavior of Plan9's `9menu` utility. it's written in go and uses fyne to provide a cross-platform gui.

## why?

this project was created because there was no version of `9menu` available for apple silicon, and a thing at work came up where a tool like this would be useful.

## usage
to use 8menu, simply run the program and provide a list of menu items as command-line arguments. each menu item should consist of a label and a command, separated by a space. 

**example:**

```bash
8menu "item 1" "echo hello" "item 2" "echo world" "item 3" "echo fyne"
```

this will create a menu with three items, labeled "item 1", "item 2", and "item 3", each of which will execute the corresponding shell command when clicked.

to exit the menu, simply click the "exit" button or close the window.
building

## download

to download a precompiled binary, click here: https://github.com/donuts-are-good/8menu/releases/latest

## build

to build 8menu from source, you will need to have go installed. once you have these dependencies installed, you can build 8menu like this:

```bash
go build
```

this will create a binary file called 8menu in the current directory.
license

8menu is distributed under the mit license. see the license file for more information.


## license
mit license 2023 donuts-are-good https://github.com/donuts-are-good
