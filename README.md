# Wordfind
If you provide wordfinder with a set of letters (this can include one letter more than once) and a word length, from its database, it will find all words with the provided length that are constructable from the set of letters.



## Usage

Use

    wordfind

to enter the programm. With "find" you can let the programm find some words for you. Use "exit" to exit the programm.

There is already a database provided (WordListDatabase.json, must be in the same directory as the executable) containing ~250k german words.
To compile your own databse, delete the WordListDatabase.json file, and use

    wordfind -b <.txt file>
   
   This will update (or create if there is no databse) the database with the words contained in the .txt file seperated by spaces or newline (there may be added more options in the future).
   To only add a few words, use
	
    wordfind -s <List of words>

To delete words, use

    wordfind -d <List of words>
    
The list of words is space separated.

All commands can be run with -v option to get a detailed output.