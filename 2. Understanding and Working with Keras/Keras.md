# Understanding and Working with keras
## Two major kinds of framework
1. Sequential API and
2. Functional API

## Major Steps to Develop Deep Learning Models With Keras
1. Declare the model. Each layer can contain one or more convolution, pooling, batch normalization,
and **activation function.**
2. Compile the model . Here you apply the **loss function** and **optimizer** before calling the compile() function on the model.
3. Fit the model with training data. Here you train the model on the test data by calling the fit() function on the model.
4. Make predictions. Here you use the model to generate predictions on new data by calling functions such as evaluate() and predict().

## Eight steps in the deep learning process in Keras
1. Load the data
2. Preprocess the data
3. Define the model
4. Compile the model
5. Fit the model
6. Evaluate the model
7. Make predictions
8. Save the model