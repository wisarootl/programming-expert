# Contact List

The [ProgrammingExpert](https://programmingexpert.io) solution to the programming project, Contact List.

## Problem Statement

For this project you'll be asked to create a Contact List program that can persistently store a list of a user’s contacts. Your program needs to create, delete, list and search for contacts. Contacts should be stored persistently, this means each time the program is executed the contacts added or deleted in previous executions are still accessible.

To store contacts persistently you should use a `.json` file. You can structure this file however you like but we have included some starting code that reads and writes the `.json` file for you. You do not need to use this starting code, it is simply here to illustrate how to use Python's `json` module.

**JSON** stands for JavaScript Object Notation and represents objects in a very similar way to Python dictionaries. To deal with `.json` files in Python we can use the `json` module as follows:

```python
import json
contacts = {"contacts": [{"first_name": "Tim", ...}]}

with open("contacts.json", "w") as f:
    json.dump(contacts, f)  # writes a Python dictionary as a json file

with open("contacts.json", "r") as f:
    contacts = json.load(f)  # loads the file as a Python dictionary

print(contacts) # {"contacts": [{"first_name": "Tim", ...}]}
```

Each contact in the contact list requires a first name and last name but all other fields are optional. The optional fields that you should prompt the user to input are the following:

- Mobile Phone Number
- Home Phone Number
- Email Address
- Address

Contacts in the contact list are unique, meaning no two contacts can have both the same first and last names. However, contacts can share a first name or a last name, just not both.

Your program should work similarly to a command line interface, where you constantly ask the user to enter one of the following commands:

- **add**: Adds a contact to the contact list.
- **delete**: Deletes a contact from the contact list using it's first and last name.
- **list**: Lists all the contacts in alphabetical order by first name.
- **search**: Searches for contacts by their first and last name.
- **q**: Quits the program and saves all of the current contacts.

When adding contacts ensure that all phone numbers that are entered are valid. For this project a phone number is valid if it is 10 digits and only contains numbers. Do the same with the email (a helper function is provided that verifies emails for you).

When searching for contacts simply display all contacts that contain the strings used to search for first and last names. For example, if your search strings are `first_name = "t"` and `last_name = "r"` then you should display the following names:

- Stan Right
- Tim Archer
- Thomas Rong

All names that contain a `t` in the first name and a `r` in the last name are displayed. You should also display all of the search results in alphabetical order.

For a more detailed example please refer to the sample program executions below and the video explanation.

### Sample Program Execution 1

```
Welcome to your contact list!
The following is a list of useable commands:
"add": Adds a contact.
"delete": Deletes a contact.
"list": Lists all contacts.
"search": Searches for a contact by name.
"q": Quits the program and saves the contact list.

Type a command: add
First Name: Tim
Last Name: Ruscica
Mobile Phone Number: 905-676-0090
Home Phone Number:
Email Address: tim@algoexpert.io
Address:
Contact Added!
Type a command: list
1. Tim Ruscica
        Mobile: 905-676-0090
        Email: tim@algoexpert.io
Type a command: add
First Name: Tim
Last Name: Ruscica
Mobile Phone Number:
Home Phone Number:
Email Address:
Address:
A contact with this name already exists.
You entered invalid information, this contact was not added.
Type a command: add
First Name: Tim
Last Name: Jones
Mobile Phone Number: 1234
Home Phone Number: 1234
Email Address:
Address:
Invalid mobile phone number.
You entered invalid information, this contact was not added.
Type a command: add
First Name: Tim
Last Name: Jones
Mobile Phone Number: 905-727-0987
Home Phone Number:
Email Address: tim@invalid
Address:
Invalid email address.
You entered invalid information, this contact was not added.
Type a command: add
First Name: Tim
Last Name: Jones
Mobile Phone Number: 905-805-9999
Home Phone Number:
Email Address: tim.jones@gmail.com
Address:
Contact Added!
Type a command: list
1. Tim Ruscica
        Mobile: 905-676-0090
        Email: tim@algoexpert.io
2. Tim Jones
        Mobile: 905-805-9999
        Email: tim.jones@gmail.com
Type a command: search
First Name: Tim
Last Name: O
Found 1 matching contacts.
1. Tim Jones
        Mobile: 905-805-9999
        Email: tim.jones@gmail.com
Type a command: q
Contacts were saved successfully.
```

## Running The Code

- To run the code first clone the repository using `git clone https://github.com/wisarootl/programming-expert.git`.
- Next `cd programming-expert/programming-project/contact-list` and execute the `main.py` file with `python main.py` or `python3 main.py`.
