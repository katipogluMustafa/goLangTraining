# Interface And Composition Design

#### Design Philosophy

* Interfaces give programs structure

* Interfaces encourage design by composition.

* Interfaces enable and enforce clean divisions between components.
    * The standardization of interfaces can set clear and consistent expectations.
* Decoupling means reducing the dependencies between components and the types they use.
    * This leads to correctness, quality and performance.
* Interfaces allow you to group concrete types by what they do.
    * Do not group types by a common DNA but by a common behaviour.
    * Everyone can work together when we focus on what we do and not who we are.     
* Interfaces help your code decouple itself from change.
    * You need to understand what could change and user interfaces to decouple.
    * Interfaces with more than one method have more than one reason to change.
    * Uncertainty about change is not a license to guess but a directive to STOP and learn more!
* You mush distinguish between code that:
    * defends against fraud vs protects against accidents.       

#### Validation:
    
Use an interface when:

* Users of the API need to provide an implementation detail.
* APIs have multiple implementations they need to maintain internally.
* Parts of a API that can change need to be identified and require decoupling.

Don't use and interface:

* For the sake of using an interface.
* To generalize an algorithm.
* When users can declare their own interfaces.
* If it's not clear how the interfaces makes the code better.