# json-visualizer
A utility for visual display of an array of data stored in a JSON file.

As an argument, the utility is passed the path to a file that contains an array of objects in the items key. Each object has the following keys:
name - student surname and first name (string type) 
group - group number (string / integer type) 
avg - average score (string / integer / real type) debt - list of debts (string / enumerated type)

An example file: data.json

Arguments:
    -dir argument specifying the path to the JSON file. (By default, the root directory is used).
