  * Here are some guidelines around interface pollution:
  
 * Use an interface:
     * When users of the API need to provide an implementation detail.
     * When API’s have multiple implementations that need to be maintained.
     * When parts of the API that can change have been identified and require decoupling.
 * Question an interface:
     * When its only purpose is for writing testable API’s (write usable API’s first).
     * When it’s not providing support for the API to decouple from change.
     * When it's not clear how the interface makes the code better.

 * Factory functions initialize concrete type should return the concrete values or pointers
    * The caller will deal with decoupling if it is needed.