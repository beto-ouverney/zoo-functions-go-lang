# Project Zoo Functions build with Go Lang

# About

This project is dedicated to the learning of Go Lang

All files in project-zoo-functions-go

# Skills used in this project:

* Readable, concise and expressive code
* Variables, Functions, Packages, etc
* Array & Slices
* Pointers, Structures, Methods
* Interface
* Git
* Data Structure and Algorithms
* Kanban

# The Project

This project is divided into questions

# Question 1

Function `getSpeciesByIds`

This function is responsible for searching animal species by id. It returns an array containing the species referring to the ids passed as a parameter, and may receive one or more ids.

# Question 2

Function `getAnimalsOlderThan`

This function, based on the name of a species and a minimum age, checks if all animals of that species have the specified minimum age.

# Question 3

Function `getEmployeeByName``

This function is responsible for searching for collaborating people by their first or last name

# Question 4

Function `getRelatedEmployees`

- if it is a manager collaborating person, it must return an array containing the names of the collaborating people it is responsible for;

Output Example: [ 'Burl Bethea', 'Ola Orloff', 'Emery Elser' ];

- if not a manager contributor, return the message "The id entered is not a manager contributor person!"

# Question 5

Function countAnimals

This function is responsible for counting the number of animals of each species.

# Question 6

Function `calculateEntry`

It should receive the visitors array and return an object with the count according to the following sort criteria:

Persons under the age of 18 are classified as children;
People aged 18 or over and under 50 are classified as adults;
People aged 50 or over are classified as older (senior) people.

# Question 7

Function `getAnimalMap`

The function is responsible for the geographic mapping of the species and their animals, and can also filter them by alphabetical order and sex.

# Question 8

Function `getSchedule`

The function is responsible for making the time information of the animals available in a query to the user, who may want to have access to the schedule of the week, a day or a specific animal.

# Question 9

Function `getOldestFromFirstSpecies

The function searches for information on the oldest animal of the first species managed by the person collaborating with the parameter.

# Question 10

Function `getEmployeesCoverage`

This function will be responsible for associating coverage information of the employees.`

Output example: 
{
  "id": "4b40a139-d4dc-4f09-822d-ec25e819a5ad",
  "fullName": "Sharonda Spry",
  "species": [ "otters", "frogs" ],
  "locations": [ "SE", "SW" ]
}

When being called with no arguments, it should return an array with the coverage of all the employees:

[
  {
    "id": "c5b83cb3-a451-49e2-ac45-ff3f54fbe7e1",
    "fullName": "Nigel Nelson",
    "species": [ "lions", "tigers" ],
    "locations": [ "NE", "NW" ],
  },
  {
    "id": "0e7b460e-acf4-4e17-bcb3-ee472265db83",
    "fullName": "Burl Bethea",
    "species": [ "lions", "tigers", "bears", "penguins" ],
    "locations": [ "NE", "NW", "NW", "SE" ],
  },
  {
    "id": "fdb2543b-5662-46a7-badc-93d960fdc0a8",
    "fullName": "Ola Orloff",
    "species": [ "otters", "frogs", "snakes", "elephants" ],
    "locations": [ "SE", "SW", "SW", "NW" ],
  },
  //[...]
];

If no person is found with the first name, last name or id, the message 'Invalid information' should be displayed



