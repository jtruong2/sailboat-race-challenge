## REQUIREMENTS

This coding exercise is an opportunity for you to show us how you break down product requirements into actual code, as well as to demonstrate code quality and style.

##### Language
Feel free to choose any language you are comfortable with. Our engineers typically work with Go, Python, and Javascript. However, your project submission is not required to be in these languages if you are not familiar with them. We value the solution more than specific tools and languages.

##### Project Setup
Please include a README file including:
- a list of dependencies of your program, as well as how to install them if applicable (we may not be experts in your chosen language and tools)
- instructions for running your program

## THE PROBLEM

##### Race Average

There is a sailboat race from Rhode Island to Bermuda. It takes several days.
`Take an array of the times` when the competitors crossed the finish line and
`convert that into the average number of minutes` to complete the race.

The race `starts on day 1 at 8:00 AM.`

We are given a `list of finish times as a string`, where each finish time is
formatted as

`hh:mm xM, DAY n`

where `hh` is exactly 2 digits giving the hour, `mm` is exactly 2 digits
giving the minute, `x` is either `'A'` or `'P'`, and `n` is a positive integer
`less than 100 with no leading zeros`. So each string has `exactly 15 or 16 characters` (depending on whether n is less than 10).

Create a program that contains an "average minutes" function. This function is given a
list of strings, `times`, and that returns the average number of minutes taken by the
competitors to complete the race. Round the `returned value to the nearest
minute, with .5 rounding up`.

##### Notes

    * `"12:00 AM, DAY d"` refers to midnight between DAY `d-1` and DAY `d`

    * `"12:00 PM, DAY d"` refers to noon on DAY `d`

##### Constraints

    * `times` contains between 1 and 50 elements inclusive

    * each element of `times` is formatted as specified above, with `hh` between
      `01` and `12` inclusive, `mm` between `00` and `59` inclusive, and `d`
      between `1` and `99` inclusive
    * each element of `times` represents a time later than the start of the race.

##### UI Design (for Frontend Engineering candidates only)
The user interface should be based on the designs given, though you may enhance functionality as you wish. At a bare minimum, the UI functionality should consist of:
* at least one user input so a user can create a list of race times
* a way to submit the time list input(s)
* a way to show the resulting average race time to a user

UI Designs will be evaluated for their consistency to the design, in addition to semantic html and general code quality.

##### Examples

Input:

    ["12:00 PM, DAY 1",
    "12:01 PM, DAY 1"]

Output:
    241

From 8:00 AM to noon is 4 hours, so we have 4 hours for one competitor, and 4
hours, 1 minute for the other. These two values average to 240.5 minutes,
which is rounded up.


Input:

    ["12:00 AM, DAY 2"]

Output:
     960

The one competitor finished in 16 hours, just at the start of DAY 2.

Input:

    ["02:00 PM, DAY 19",
    "02:00 PM, DAY 20",
    "01:58 PM, DAY 20"]

Output:

27239