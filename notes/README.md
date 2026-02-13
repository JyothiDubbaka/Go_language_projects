# Notes Tool
## Overview

The Notes Tool is a simple command-line application that allows users to manage short, single-line notes organized in collections. Each collection is stored in a separate text file, and notes persist between runs of the program. Users can create, view, add, and delete notes interactively from the terminal.

The tool is designed to help users organize their ideas or tasks in a lightweight and easy-to-use manner.

## Features

- Create or open a collection of notes.

- Display all notes in a collection.

- Add a new note to a collection.

- Delete an existing note from a collection.

- Persistent storage — notes are saved in plain text files.

- Interactive menu with clear options.

- Built-in help message for usage instructions.

## Usage

1. Load the tool using git
```shell
git clone https://gitea.kood.tech/malikquamrussamawat/notes 
```

2. Run the tool in the command-line using command:
```shell
./notestool [COLLECTION_NAME]
```

3. COLLECTION_NAME is the name of the notes collection. If the collection does not exist, a new text file will be created with a name COLLECTION_NAME.txt

4. If no argument is provided or help is given, the tool displays a help message:
```shell
./notestool
Usage: ./todotool [TAG]

./notesJtool help
Usage: ./todotool [TAG]
```

## Example Workflow

1. Open or create a collection
```shell
./notestool testtag

Welcome to the notes tool Folks! 🤝
```

2. Menu options
```shell
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1
```

3. Show Notes
Select 1 to show notes
```shell
Notes:
001 - note one
002 - note two
```

4. Add a Note
Select 2 to add a new note
```shell
Enter the note text:
note three
```
After adding, select 1 again to see updated notes:
```shell
Notes:
001 - note one
002 - note two
003 - note three
```

5. Delete a Note
Select 3 to delete a note
```shell
Enter the number of note to remove or 0 to cancel:
2
```
After deleting, select 1 to view remaining notes:
```shell
Notes:
001 - note one
002 - note three
```

6. Exit the tool
Select 4 to exit the program.
```shell
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
```
## Building the executable 
### Initialising package
```shell
go mod init notes
```
### Building the executable
```shell
go build -o notestool 
```
```-o notestool``` defines the executable name and hence the application could be run using ```./notestool```.

## Data Storage

- Each collection is stored in a plain text file with the same name as the collection.

- Each note is stored on a separate line.

Example content of testtag after some operations:
note one
note two
note three

Notes persist between runs of the program.

The tool uses the following file handling strategies:
    
    os.O_CREATE | os.O_RDONLY → safely open or create a file for reading.

    os.O_TRUNC | os.O_WRONLY | os.O_CREATE → overwrite file when saving updated notes.

    bufio → used to read/write files efficiently line by line.

## Allowed Packages

bufio — buffered reading/writing of files and input

fmt — formatted input/output

os — file and system operations

strconv — converting string input to integers

strings — string manipulation

## Summary

The Notes Tool provides a simple and interactive way to manage single-line notes in collections.
It supports:

- Persistent storage

- Adding, viewing, and deleting notes

- Easy-to-use command-line interface

- This tool is suitable for quick note-taking, idea tracking, or lightweight task management.
