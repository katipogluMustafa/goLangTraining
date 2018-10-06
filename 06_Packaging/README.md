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

## Package Oriented Design

Let's talk about project structure that Bill Kennedy use on all of his projects.

### Kit Project

* I really believe that every company or at least every team should have a kit project.
    * A kit project is a set of foundational packages or APIs that every application you're building should use.
    * So I would wanna see things like 
        * log packages if you're not using the standard library.
        * config packages, if you're not using some third part stuff.
        * web frameworks you're building.
    * Whatever it is, I also like to see that laid out in the root of source tree.
        * Probably don't wanna see, package inside a packages here.
    * I also want these packages as decoupled as possible.
        * In other words, I don't want log importing config, or config importing log.
        * Really these packages shouldn't even be logging at all.
            * Because if you choose to do some sort of logging, then you're setting policy.
            * I'm also not a big fan of logging interface, I think that's a big big mistake.
            * I mean, standard library doesn't log, these foundational packages shouldn't log either.
    * If these foundational packages have goroutines or paths of execution where events are happening that you wanna log, then I think you should ask for a handler function and the user can implement whatever they want in the handler function and you just call during those events.
        * So you don't need interfaces all the time.
        * Asking for a function like a handler function really can help streamline everything and then get you to that logging that you need.
 * That's the kit, it should be in it's own repo.Hopefully something called kit.
                      
```
Kit                     Application

├── CONTRIBUTORS        ├── cmd/
├── LICENSE             ├── internal/
├── README.md           │   └── platform/
├── cfg/                └── vendor/
├── examples/
├── log/
├── pool/
├── tcp/
├── timezone/
├── udp/
└── web/
```    

### Application Project

* Every project that you work on is what I call an application project.
    * Application project can have multiple binaries in it. This isn't a one to one.
        * In fact less is always more.
    * A large team of people can work in an application project if package oriented design is being implemented properly without stepping on each other.
    * I want you to know that there are four folders in an application project.
        * cmd 
        * internal
        * internal/platform
        * vendor    
    * This indicates us where packages should go, because each one has very specific purpose.

#### Vendor Folder

We're gonna be using vendor folder DEP, D-E-P.

* That's is the recommended tool for dealing with reproducibility and vendoring.
* DEP tool basically, on a new project you'l say "dep init", and then just "dep ensure"
    * And then dep is gonna go and look at all of the third party packages, including kit(kit will be third party package).And you'll maintain all of the source code and the versions of that.
        * You need this, please own all of the source code unless you've got project as big as kubernetes , you should be able to own all of the source code and the dependencies that you're working with.
        * Use dep[if you are reading this, research VGO, if it includes this tool in itself, you should be working with it but as of 2018 october, it still has at least 1 year of time to be ready for production.]

### Cmd (Command) 

Command is where the applications are, where the binaries that we're building, it can be one to many. 

* [On this service project](https://github.com/ardanlabs/service), which is implementing a crud-based web service, you can see a folder called "crud" and you can see a "main.go". "Main.go" is that entry-point and it's where we're building. 
    * That means that "crud" will be the name of the binary. 
    
    * This is important. If we have multiple binaries we're building they could be under command or in this case since I'm using a "sidecar" architecture, you could see that we have two other binaries, one called "metrics", one called "tracer". These are micro-services that help deal with tracing and metrics for this service.
    
    * These are three different binaries that we build. Three different "main.go". 
    
    * Now, packages that are defined inside of let's say "crud" are very application specific and they're there to help support start-up, shut-down, maybe some small routing like in a web service. But, **not a lot of business logic**, more **presentational logic**, like in a web service. Taking a request, doing a response. 
    
    * Now, what's nice is at the application layer, right at the command layer, these packages can contain things if you really need to. You can get away with packages that contain at this level because they have really no level of reusability. 
        * They're specific for this app and nothing else. 
        * You also could have test folders here or packages for integration tests. That can help and work with you as well. 
        
    * The application layer is where you have a lot of ability to set policy. You could do some containment like we're doing with handlers. And you're going to have multiple binaries that you're building.        
    
### Internal Folder

Packages that need to be imported by multiple programs within the project belong inside the internal/ folder. One benefit of using the name internal/ is that the project gets an extra level of protection from the compiler. No package outside of this project can import packages from inside of internal/. These packages are therefore internal to this project only.

### Internal Platform Folder

Packages that are foundational but specific to the project belong in the internal/platform/ folder. These would be packages that provide support for things like databases, authentication or even marshaling.

