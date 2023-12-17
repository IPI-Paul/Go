# Derek Banas Go Tutorials

[Golang Tutorial Go Full Course](https://youtu.be/YzLrWHZa-Kc?si=4kqRN-8Pj-9Jp580)

## Course Work

 - Once again, Derek goes over a whole lot of examples and I have put them into 42 functions.
 - Unlike Rust where you can completely contain structs, implementations, traits and functions within a function, I found that Go only allows use of closures instead of functions within functions and structs within functions could not have functions assigned to them.
 - However, I made good use of assigning functions to a struct and using reflect's MethodByName to call each example as per the user's command line inputs.
 - I have not yet figured out how to import the external apps, as the previous course only showed me how to create a binary, but attempting to create a library failed. So, I have called the external apps using the exec.Command function which requires the current folder structure.
 - Unfortunately, Derek does not cover how o create a binary or pkg.
 - The web app is something I can use on a daily basis and unlike Angular, runs inside a python pyqt5 browser.