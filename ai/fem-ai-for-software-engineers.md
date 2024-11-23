# AI for Software Engineers

> Notes from [this workshop](https://frontendmasters.com/workshops/engineering-and-ai/)

## Concepts

- In the beginning of the workshop, we have created a basic _converted_ that, given some parameters, like amount of money and the age of the account, returned information whether the user should be given a refund or not.

  - **This _converter_ is called a _model_**. There are different types of models. In our case, we created a **_decision tree model_**.

- The goal is to create a _converter_ that is the most accurate against our sample data.

  - This is quite tricky, because we have to ensure the sample data is representative of a real-world population (whatever that might be).

- Deploying any model to "production" is a multi-step process that often involves a lot of data-science related work and experimentation.

  - One example would be _pre-processing the data_ to see how the _converter_ behaves.

    - Perhaps we do not care about cents and look only at the amount of dollars as one of the parameters to decide whether to grant a refund or not.

- You can _tune the model_ after the training phase.

  - **_Tuning_ the model is quite important. It allows you to prevent "drift" where model starts to product inaccurate results**.

    - Remember that the circumstances under which the model underwent the training might be different that the current reality.

- **If we can create _converters_ based on numbers, we can also convert other things to numbers and create _converters_ for those things**.

  - Images are, set of pixels, so they can be converted to values.

  - Words are a bit harder to convert to words, but it could be done. Think _embeddings_.

## Neural networks

- As an example, we have developed a grid with filled with "multipliers".

  - If you multiply the "image grid" with the "multipliers grid" and all numbers together, you will get either -1 or 1.

    - 1 indicates a smile, -1 indicates a sad face.

- **Running the _converter_ on some input, and getting the result back** is called **inference**.

  - So, in our case, multiplying the "multipliers" grid with the grid representing an image, and getting the result back, would be called _inference_.

- **Out "multipliers" grid contains "weights"**. The grid itself is **called a "layer"**.

### Figuring out the weights

- First we started with our grid filled with zeros.

  - When, we tweaked the "significant" positions one by one.

  - Then, we applied those changes all at once.

    - Surprisingly, the net-result was positive. We got more accurate results! Note that **we did not get 100% accurate**.

      - This was surprising to me, because I would not think that, a _sum_ of changes in the "right direction" would point us to the right direction.

    - **This technique is called _gradient descent_**.

Finished -34:29 part 4
