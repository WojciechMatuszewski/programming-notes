# Rust: The Complete Developer's Guide

> Learning from [this course](https://www.udemy.com/course/rust-the-complete-developers-guide/).

## Ownership and Borrowing

- The _variables_ in Rust can sometimes be referred to as _bindings_.

- **Rust wants to avoid "unexpected updates"**.

  - That is why you have the ownership and lifetimes. In the end, it makes your programs less bug-prone.

- There are **two types of _"moves"_**.

  - The "full" move.

    - This is where the whole structure was moved.

  - The partial move.

    - This occurs when some property on a structure was moved. Rust will not let you use the whole structure again.

- You can create **as many _read-only_ references to a given value as you wish**.

  - Makes sense. If the reference is _read-only_, there is no risk of mutating it.

- You **can't move a value if there are references to that value or properties on the value**.

  ```rust
      fn main() {
      let mut bank = Bank::new();
      let account = Account::new(1, String::from("me"));

      let holder_ref = &account.holder; // reference to property on the struct

      bank.accounts.push(account); // attempt to move here. Error!

      println!("{}", holder_ref);
      }
  ```

- **There can only be one mutable reference** to a given value.

  - You **can't use both mutable and non-mutable references to a given value at the same time**.

  ```rust
    fn main() {
      let mut bank = Bank::new();
      let bank_ref = &bank; // non-mutable reference

      add_account(&mut bank); // mutable reference

      println!("bank = {:?}", bank_ref); // error
  }
  ```

- **Some values may appear to break the rules of ownership and borrowing**.

  - Those values, mostly primitives, **are using _copy_ instead of _move_ semantics**.

  ```rust
  fn main() {
    let id = 1;
    let other_num = id; // Copy instead of move

    println!("{:#?} {:#?}", id, other_num);
  }
  ```

### Temporary refs

```rust
fn print_account(account: &Account) {}

fn main() {
  let mut bank = Bank::new();

  let account = Account::new();
  let account_ref = &account; // This is NOT a temporary ref.
  print_account(account_ref)

  // account_ref still exists here
  bank.accounts.push(account); // Error! Both ref and the value exist. You can't move the value while ref to the value exist.
}
```

Contrast the above with the following.

```rust
fn print_account(account: &Account) {}

fn main() {
  let mut bank = Bank::new();

  let account = Account::new();
  print_account(&account)

  // The `&account` does not exist here.

  bank.accounts.push(account); // Works

}
```

## Lifetimes

- They answer the question: _"how long an owner/reference exists"_.

- To understand the _lifetimes_ it is imperative to understand when a given value is _dropped_ by Rust compiler.

```rust
fn print_account(account: Account) {
  // The `account` argument is dropped after this function is exited
}

fn main() {
  let account = Account::new();

  print_account(account);

  // You can't use `account` variable here.
  // It is no longer in memory.
}
```
