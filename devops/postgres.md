# Postgres

Learning about the Postgres database!

## Row Level Security (RLS)

> Based on [this talk](https://www.youtube.com/watch?v=vZT1Qx2xUCo)

- **Authnz for Postgres**

  - Authentication: _Is this user/role allowed to access the database?_

  - Authorization: _Now that they are in, what specifically are they allowed to access?_

- **You can think of the RLS policies as "implicit" `where` clauses**.

  - Instead of writing `select * from profiles`, Postgres would automatically wire this statement through the relevant policy which might apply some constraints, like `select * from profiles where user_id = 123`.

- **With RLS comes a risk of introducing performance issues**.

  - Keep in mind that those policies will run when you perform a query. **If you are not mindful how you write them, you might slow your database down**.

    - [See this part of the video](https://youtu.be/vZT1Qx2xUCo?t=793) for relevant tips.

My personal take is that, this approach, is like putting logic inside your SQL statements. I'm unsure if I'm a fan. We all know that, in theory, _stored procedures_ are a great idea, but in reality they bring a world of pain. I'm yet to use RLS in any kind of application, so my opinion is not formed on facts, but rather a gut instinct.
