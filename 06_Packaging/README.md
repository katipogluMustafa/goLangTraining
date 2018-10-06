# Package Oriented Design

## Language Mechanics

Package Oriented Design gives you an engineering decision about where packages go and it's going to help you identify how your go project should be structured and it's gonna really improve communication  between you and your team members.

The whole idea here is that we want clean package design and project architecture that is not random but very predictable.

If you don't have a package oriented design on your project structure and your packages you really going to have lots of problems.Maybe not today, maybe not a month from now but you will definitely in years.

---

* There is no real concept of subpackages just because a folder sitting inside another folder doesn't mean that this package is a subpackage.

* All packages built and compiled and then laid out.

* From the linkers point of view all packages are at the same level.
    * So we wanna leverage the hierarchy to give us some indications about relationships.
    * But from the compilers point of view there are not subpackages, there no relationships, all packages are at th same level.

* Another interesting thing is that two packages can not cross import each other.

* Language Mechanics are kind of few, but they do they do put some interesting constraints on how we define projects, lay out projects , how we code and architect and design our source code.
    * It's no more monolithic applications with folders, it is a project with lots of individual static libraries that we call packages that eventually are going to import and be bound together.

## History

In an interview given to Brian Kernighan by Mihai Budiu in the year 2000, Brian was asked the following question:

**_“Can you tell us about the worse features of C, from your point of view”?_**

This was Brian’s response:

**_“I think that the real problem with C is that it doesn’t give you enough mechanisms for structuring really big programs, for creating "firewalls" within programs so you can keep the various pieces apart. It’s not that you can’t do all of these things, that you can’t simulate object-oriented programming or other methodology you want in C. You can simulate it, but the compiler, the language itself isn’t giving you any help.”_**

