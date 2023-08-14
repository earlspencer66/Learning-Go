# Learning-Go
I'm learning Golang - one step at a time
## Chapter 1

### Basics of TensorFlow
* *Tensor* - a mathematical model, a generalization of vectors, scalars and matrices.
* Can be considered to be an n-dimensonal array.
* *tensorFlow program* - construction phase and execution phase.
* *Construction phase* - assembles a graph that has nodes (operations) and edges (tensors)
* *Exceution phase* - session to execute the operations in the graph.
* A *session* encapsulates the control and state of the TensorFlow runtime
* Object to create a session object
* ```python
    sess = tf.Session()
    ```
* *Variables* enable us to enter trainable parameters into the graph
* >Constants are initialized when you call tf.constant, and their values can never change. By contrast, variables are not initialized when you call tf.Variable. To initialize all the variables in a TensorFlow program, you must explicitly call a special operation as follows.
  ```python
  sess.run(tf.global_variables_initializer())

  # OR

  sess.run(tf.compat.v1.global_variables_initializer())
  ```
  ## Creating Tensors
  
  * Activation functions - from analysis of how a neuron works in the human brain.
  * Neutron becomes active beyond a certain potential i.e. the *activation threshhold*
  * Most populat activation functions are (1) hyperbolic tangent (tanh), (2) ReLu and (3) ELU.
  * Tangent Hypebolic and Sigmoid
  * Tangent Hyperbolic
  * ```latex
      \begin{equation}
          tanh x = \frac{e^x - e^{-x}}{e^x+e^{-x}}
      \end{equation}
      Output will always lie between -1 and 1
       ```
  * Sigmoid
  ```latex
    \begin{equation}
        sigmoid x = \frac{1}{1+e^{-x}}
    \end{equation}
      Output will always lie between 0 and 1
  ```