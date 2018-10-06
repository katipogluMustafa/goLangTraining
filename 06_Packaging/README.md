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

## Design Guidelines

* KeyWords: 
    * **Purposeful**
    * **Usable**
    * **Portable**

* **To be purposeful, packages must provide, not contain**.
    * Packages must be named with the intent to describe what it provides.
    * Packages must not become a dumping ground of disparate concerns.
* **To be usable, packages must be designed with the user as their focus**.
    * Packages must be intuitive and simple to use.
    * Packages must respect their impact on resources and performance.
    * Packages must protect the user’s application from cascading changes.
    * Packages must prevent the need for type assertions to the concrete.    
    * Packages must reduce, minimize and simplify its code base.
* **To be portable, packages must be designed with reusability in mind**.   
    * Packages must aspire for the highest level of portability. 
    * Packages must reduce setting policy when it’s reasonable and practical.
    * Packages must not become a single point of dependency.
---

* **To be purposeful, packages must provide, not contain**.
    * Every package must have an API, a clear API on what it provides to user. And if its not clear, then question the packages existence.
    * The name of the package should describe the intent of what it provides.
        * If you struggling the name a package, you probably code smell there.
        * Packages like fmt, net, http, os, that's clear, we know what they provide.
        * Packages like util, helper, common these are packages that contain code and these are gonna cause you a lot of problems because they're gonna be a dumping ground of code, they're gonna be single point of dependency, and eventually you're going to crumble on top.
        * And I promise you this, If you have a package called models, a package that is common set of types, your project has already failed.
            * In monolithic applications you can have that type of encapsulation from a type perspective but not here in go.
            * Types are an artifact to move data across the program boundaries. They cannot be API in and of themselves.
            * Avoid packages like models.
    * Every purposeful package, every package that has real purpose, has its own type system, even if you have to duplicate types across multiple packages that's going to be much better.
        * Remember concrete data solves the problem.
        * We can leverage interfaces later on to decouple our APIs and to accept everybody's concrete data.
                 
---

Usability is more of an art.

* Remember also that usability means that we're designing APIs and packages to reduce or minimize fraud and misuse.
    * Precision is everything.
        * We don't want packages and APIs that are vague

---  

The more decoupled, the more reusable a package is the better it's going to be in your ecosystem.


* The more policy a package has the less reusable it is.
    * When I talk about policy what I mean is if a package is making decisions about how we log, how we do configuration, how we do things, then only other applications that wanna do those things the same can use that package.
    